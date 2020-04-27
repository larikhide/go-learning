package main

import "fmt"

var x string = "Hello Tello" /* если задать переменную за функцией main, то к ней можно обращаться через другие функции */

func main() {
	fmt.Println(x)
}

func f() {
	fmt.Println(x)
}
