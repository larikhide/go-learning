package main

import "fmt"

const koefFootsToMeters = 0.3048

func main() {
	fmt.Print("Enter a foots: ")
	var foots float64
	fmt.Scanf("%f", &foots)
	meters := foots * koefFootsToMeters
	fmt.Println(meters)
}
