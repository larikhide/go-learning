package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var number int
	fmt.Scanf("%i", &number)
	switch number {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("unknown number")
	}
}
