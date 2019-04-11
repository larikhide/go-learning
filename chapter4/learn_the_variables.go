package main

import "fmt"

func main() {
	fmt.Print("Enter a foots: ")
	var foots float64
	fmt.Scanf("%f", &foots)

	meters := foots * 0.3048

	fmt.Println("foots in metres are ", meters)
}
