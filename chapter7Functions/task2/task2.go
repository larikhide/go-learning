package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxElement(2, 4, 98, -43))
}

func maxElement(list ...int) int {
	maxNum := math.MinInt32
	for _, value := range list {
		if maxNum < value {
			maxNum = value
		}
	}
	return maxNum
}
