package main

import "fmt"
import "D:/skripts/Study-to-be-a-golang-programmer/chapter11Packages/math"

func main() {
	xs := []float64{}
	avg := math.Average(xs)
	fmt.Println(avg)
}
