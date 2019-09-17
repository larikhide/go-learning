package main

import "fmt"

func binarySearch(list []int, item int) int {
	low := 0              //индекс первого элемента среза
	high := len(list) - 1 //индекс последнего элемента
	var mid int
	for low <= high {
		mid = (low + high) / 2
		switch {
		case list[mid] == item:
			return list[mid]
		case list[mid] > item:
			high = mid - 1
		case list[mid] < item:
			low = mid + 1
		}
	}
	return list[mid]
}

func main() {
	list := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&list[i])
	}
	fmt.Println(binarySearch(list, 4))
}
