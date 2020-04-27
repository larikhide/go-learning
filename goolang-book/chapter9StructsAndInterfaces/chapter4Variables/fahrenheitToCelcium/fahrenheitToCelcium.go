package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celcium := (fahrenheit - 32) * 5 / 9

	fmt.Println(celcium)
}
