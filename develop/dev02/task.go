package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	errBadInput = errors.New("неверная строка")
)

func unpack(text string) (string, error) {
	var letter, answer string
	if text == " " {
		answer = " "
	} else {
		for i, v := range text {
			if i == 0 && v >= 48 && v <= 57 {
				answer = ""
				return answer, errBadInput

			} else if n, err := strconv.Atoi(string(v)); err == nil {
				if _, err := strconv.Atoi(string(text[i-1])); err != nil {
					letter = string(text[i-1])
				}
				for j := 0; j < n-1; j++ {
					answer += letter
				}
			} else {
				answer += string(v)
			}
		}
	}
	return answer, nil

}

func main() {

	b := " 10"
	a, err := unpack(b)
	fmt.Println(a, err)
}
