package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	width, height float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r *Rectangle) area() float64 {
	return r.width * r.height
}
func (r *Rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func measure(S Shape) {
	fmt.Println(S)
	fmt.Println(S.area())
	fmt.Println(S.perimeter())
}
func main() {
	c := Circle{radius: 6}
	r := Rectangle{width: 2, height: 4}
	measure(c)
	measure(r)

}
