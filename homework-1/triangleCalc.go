package main
import (
	"fmt"
	"math"
)

func main(){
	var (
		a, b float64 = 3, 4
		area, perimeter, hypotenyse float64
	)
	//find area of triangle
	area = a * b / 0.5

	//find hypotenyse
	hypotenyse = math.Hypot(a, b)

	//find perimeter
	perimeter = a + b + hypotenyse

	fmt.Printf("perimeter is %f\nhypotenyse is %f\narea is %f\n", perimeter,hypotenyse, area)
}
