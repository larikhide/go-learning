package main

import "fmt"

func half(num int) (int, bool) {
	r := num / 2
	return r, num%2 == 0
}

func main() {
	fmt.Println(half(2))
	fmt.Println(half(43))
}
