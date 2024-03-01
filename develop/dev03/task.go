package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var errOutOfRange = errors.New("Ошибка с выбором столбца сортировки")

func getUnique(rows []string) []string {
	set := make(map[string]struct{})
	res := []string{}
	for _, r := range rows {
		if _, ok := set[r]; !ok {
			res = append(res, r)
			set[r] = struct{}{}
		}
	}
	return res
}

func getValue(str string, col int, sep string) (string, error) {
	values := strings.Split(str, sep)
	if col >= len(values) || col < 0 {
		return "", errOutOfRange
	}
	return values[col], nil
}

func fileReader(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	storage := []string{}
	for scanner.Scan() {
		storage = append(storage, scanner.Text())
	}

	defer file.Close()
	return storage

}

type compareFunc func(a, b string) (bool, error)

func compareAsStrings(a, b string) (bool, error) {
	return a < b, nil
}

func compareAsNums(a, b string) (bool, error) {
	v1, err := strconv.Atoi(a)
	if err != nil {
		return false, err
	}

	v2, err := strconv.Atoi(b)
	if err != nil {
		return false, err
	}

	return v1 < v2, nil
}

func fileWriter(fileName string, rows []string) error {

	// Соединяем все строки в одну строку
	str := strings.Join(rows, "")

	// Преобразуем строку в срез байт
	byteSlice := []byte(str)

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Errorf("Ошибка при открытие файла")
	}
	defer file.Close()

	_, err = file.Write(byteSlice)
	if err != nil {
		return errors.Errorf("Ошибка при записи в файл")
	}
	return nil

}

func main() {
	var fileName string

	k := flag.Int("k", 1, "")
	r := flag.Bool("r", false, "")
	n := flag.Bool("n", false, "")
	u := flag.Bool("u", false, "")

	flag.Parse()

	rows := fileReader(flag.Arg(0))

	var compareStrategy = compareAsStrings
	if *n {
		compareStrategy = compareAsNums
	}

	compareFunction := func(i, j int) bool {
		r1, err := getValue(rows[i], *k-1, " ")
		if err != nil {
			log.Fatal(err.Error())
		}

		r2, err := getValue(rows[j], *k-1, " ")
		if err != nil {
			log.Fatal(err.Error())
		}

		res, err := compareStrategy(r1, r2)
		if err != nil {
			log.Fatal(err.Error())
		}

		return res == !*r

	}
	sort.Slice(rows, compareFunction)

	if *u {
		rows = getUnique(rows)
	}

	fmt.Println(rows)

	result := fileWriter(fileName, rows)

	if result == nil {
		fmt.Println("Файл отсортирован")
	}
}
