package main

import "fmt"

func main() {
	naturals := make(chan int, 10)
	squares := make(chan int)

	//генерируются простые числа до 10 и отправляются в канал
	go func() {
		defer close(naturals)
		for x := 2; x <= cap(naturals); x++ {
			// отправление в буфферизированный канал
			naturals <- x
		}

	}()

	go func() {
		defer close(squares)
		for {
			// получение значения из канала и присвоение его в  х
			// и если канал Naturals закрылся, то завершить цикл и закрыть канал
			x, ok := <-naturals
			if !ok {
				break
			}
			// отправить в канал квадрат
			squares <- x * x
		}

	}()

	for x := range squares {
		fmt.Println(x)
	}
}
