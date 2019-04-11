package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x) /* можно сделать так */

	var y string
	y = "Hello World"
	fmt.Println(y) /* а можно так */
}

func test() {
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
}
