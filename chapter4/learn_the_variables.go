package main

import "fmt"

func main() {
	fmt.Print("Enter a degree: ")
	var fharenheit float64
	fmt.Scanf("%f", &fharenheit)

	celcium := (fharenheit - 32) * 5 / 9

	fmt.Println("temperature in celcium is ", celcium)
}
