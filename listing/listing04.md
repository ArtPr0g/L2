Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1
2
3
4
5
6
7
8
9
deadlock
```
В конце будет дедлок так как мы не закрыли канал после окончания работы сс ним и main будет ждать значение из этого канала. В связи с тем что у нас работает 1 горутина, которая ждет данные из канала произайдет deadlock.
