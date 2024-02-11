Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
При сравнение статистического интерфейса тип  error  со значением  nil бдует поулчен такой же результат к и в предыдущем заданием, т.е. data == nil, itab != nil. Поэтому на выхоже будет получен error.
