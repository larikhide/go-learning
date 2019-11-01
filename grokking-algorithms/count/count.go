package main

import (
	"fmt"
)

func count(arr []int32) int {
	if len(arr) == 0 { //base case
		return 0
	} //recursive case
	return 1 + count(arr[1:])
}

func main() {
	var (
		n  int
		el int32
	)
	fmt.Scanf("%d", &n)
	arr := make([]int32, n)
	for i := range arr {
		fmt.Scan(&el)
		arr[i] = el
	}
	fmt.Println(count(arr))
}
