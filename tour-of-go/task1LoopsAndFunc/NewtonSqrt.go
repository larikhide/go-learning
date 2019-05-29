package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	d := 0.0 // переменная для оценки точности
	for i := 0; i < 10000; i++ {
		z = z - (z*z-x)/(2*z)
		if math.Abs(d-z) < 0.0001 && i != 0 { //оценка точности
			break
		}
		d = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2.214144))
}
