package main

import (
	"fmt"
	"math"
)

func main() {
	x := []int{
		48, 96, 86, 68, // найти наименьшее число массива
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	minNum := math.MinInt32
	for _, value := range x {
		if minNum > value || minNum == math.MinInt32 {
			minNum = value
		}
	}

	fmt.Println(minNum)
}
