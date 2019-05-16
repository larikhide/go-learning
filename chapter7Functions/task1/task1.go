package main

import "fmt"

func half(digit int) (int, bool) {
	r := digit / 2
	return r, r%2 == 0
}

func main() {
	fmt.Println(half(4))
	fmt.Println(half(4))
}
