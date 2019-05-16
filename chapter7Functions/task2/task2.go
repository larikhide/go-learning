package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxElement(2, 4, 98, -43))
}

func maxElement(list ...int) int {
	maxNum := math.MaxInt32
	for _, value := range list {
		if maxNum < value || maxNum == math.MaxInt32 {
			maxNum = value
		}
	}
	return maxNum
}
