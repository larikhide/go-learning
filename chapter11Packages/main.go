package main

import "fmt"
import "D:/skripts/Study-to-be-a-golang-programmer/chapter11Packages/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
