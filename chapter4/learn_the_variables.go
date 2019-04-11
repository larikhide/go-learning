package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	fmt.Print("Enter a foots: ") /* можно сделать так */

	var y string
	x = "Hello World"
	fmt.Println(x) /* а можно так */

	var foots float64
	fmt.Scanf("%f", &foots)

	meters := foots * 0.3048

	fmt.Println("foots in metres are ", meters)
}
