package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	fields := strings.Fields(s)
	for _, v := range fields {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
