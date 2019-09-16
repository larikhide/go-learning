package main

import "fmt"

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1
	var guess, mid int
	for low != high {
		mid = (low + high) / 2
		guess = list[mid]
		switch {
		case guess == item:
			return guess
		case guess > item:
			high = guess
		case guess < item:
			low = guess
		default:
			return item
		}
	}
	return guess
}

func main() {
	list := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&list[i])
	}
	fmt.Println(binarySearch(list, 4))
}
