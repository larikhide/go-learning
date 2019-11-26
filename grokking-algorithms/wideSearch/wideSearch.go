package main

import "fmt"

func main() {
	var nodes = map[int]int{
		1: 23,
		2: 36,
		7: 88,
		9: 10,
	}
	quene(3, nodes)
}

func quene(n int, list map[int]int) map[int]bool{
	row := map[int]bool {}
	if n, ok := list[n]; ok {
		row = append(row, n) //говно
}
