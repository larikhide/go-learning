package main

import "fmt"

func main() {

	totalScore := []int32{0, 0}
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}
	for i := 0; i < len(a); i++ {

		switch {
		case a[i] > b[i]:
			totalScore[0]++
		case a[i] < b[i]:
			totalScore[1]++
		}
	}

	fmt.Println(totalScore)
}
