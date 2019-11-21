package main

import "fmt"

/* func wideSearch (map[string]string) string {
	for _,
} */

func main() {
	persons := map[string]map[int]string{
		"you": map[int]string{
			0: "Alice",
			1: "Bob",
			2: "Claire",
		},
		"Bob": map[int]string{
			0: "Anuj",
			1: "Peggy",
		},
		"Alice": map[int]string{
			0: "Peggy",
		},
		"Claire": map[int]string{
			0: "Thom",
			1: "Jonny",
		},
	}

	fmt.Println(persons)
}
