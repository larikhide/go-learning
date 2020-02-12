package main

import "fmt"

func main() {
	var rate, deposit float64

	fmt.Println("insert percent rate")
	fmt.Scanf("%f", &rate)

	fmt.Println("insert deposit amount")
	fmt.Scanf("%f", &deposit)

	for i := 1; i <= 5; i++ {
		deposit += deposit * rate / 100
	}

	fmt.Printf("your sum after 5 years is %f\n", deposit)
}
