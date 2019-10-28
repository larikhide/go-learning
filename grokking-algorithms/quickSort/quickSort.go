package main

import (
	"fmt"
)

func quickSort(arr []int32) []int32 {
	var less, greater []int32
	for i := range arr {
		pivot := len(arr) / 2
		switch {
		/* case arr[i] == arr[pivot]:
		return arr */
		case arr[i] < arr[pivot]:
			less[0] = arr[i]
		case arr[i] >= arr[pivot]:
			greater[0] = arr[i]
		default:
			return arr
		}
	}
	arr = append(less, greater...)
	return arr
}

func main() {
	var (
		n  int
		el int32
	)
	fmt.Scan(&n)
	arr := make([]int32, n)
	for i := range arr {
		fmt.Scan(&el)
		arr[i] = el
	}
	fmt.Print(quickSort(arr))
}
