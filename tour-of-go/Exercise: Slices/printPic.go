package main

import (
	"golang.org/x/tour/pic"
	//"math"
)

func Pic(dx, dy int) [][]uint8 {
	dp := make([][]uint8, dy)

	for x := range dp {
		dp[x] = make([]uint8, dx)
		for y := range dp[x] {
			dp[x][y] = uint8(x*x + y*y + x*y/2*3*4)
		}
	}
	return dp
}

func main() {
	pic.Show(Pic)
}
