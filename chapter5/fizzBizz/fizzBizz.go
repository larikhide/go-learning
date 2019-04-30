package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		switch number {
		case number%3 == 0:
			fmt.Println("Fizz")
		case number%5 == 0:
			fmt.Println("Bizz")
		case number%15 == 0:
			fmt.Println("FizzBizz")
		default:
			fmt.Println(number)
		}
	}

}
