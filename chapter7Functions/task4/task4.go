package main

import "fmt"

func fibonacci(element int) int {
	if element < 2 {
		return element
	}
	return fibonacci(element-1) + fibonacci(element-2)
}
func main() {
	fmt.Println(fibonacci(2))
}
