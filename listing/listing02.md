Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
```

В функции test() defer работал по правилу - Отложенные функции могут читать и присваивать именнованные возвращаемые значения возвращающей функции, т.е. defer увеличил значение на 1.
А в функции anotherTest() он этго сделать не сможет, т.к. не будет иметь доступа к переменной x.
