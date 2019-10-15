package main

import (
	"fmt"
)

func sum(arr []int32) (total int32) {
	for i := range arr {
		total += arr[i]
	}
	return
}

/* func sumRec(arr []int32) (total int32) { //попытка через рекурсию
	if len(arr) == 0 { //базовый случай
		return total
	} else {
		sumRec(arr) + arr[] //рекурсивный
	}
} */

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
	fmt.Println(sum(arr))
}
