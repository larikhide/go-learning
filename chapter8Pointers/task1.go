package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
func main() {
	x, y := 1, 2

	swap(&x, &y)
	fmt.Println("x=", x, "y=", y)

}
