package main

import "fmt"

var x string = "Hello Tello" /* если задать переменную за функцией main, то к ней можно обращаться через другие функции */

const g = "Tello Hotello" /* константу нельзя менять по ходу выполнении программы, иначе будет ошибка */

func main() {
	fmt.Println(g)
}

func f() {
	fmt.Println(x)
}
