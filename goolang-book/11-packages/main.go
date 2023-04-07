package main

import (
	"fmt"

	"github.com/larikhide/go-learning/goolang-book/11-packages/math"
)

func main() {
	xs := []float64{}
	avg := math.Average(xs)
	fmt.Println(avg)
}
