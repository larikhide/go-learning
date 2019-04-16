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
	var z string
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z) /* можно так */

	var c string
	c = "first"
	fmt.Println(c)
	c = c + "second"
	fmt.Println(c) /* но так пизже */
}
