package main

import (
	"fmt"
)

// Complete the diagonalDifference function below.
func main() {
	var n, i, j int32
	// шаблон для заполнения двумерного массива длиной n
	fmt.Scan(&n)
	var arr [][]int32
	for i = 0; i < n; i++ {
		arr = append(arr, []int32{})
		for j = 0; j < n; j++ {
			arr[i] = append(arr[i], 0)
			fmt.Scan(&arr[i][j])
		}
	}
	fmt.Println(arr)
}
