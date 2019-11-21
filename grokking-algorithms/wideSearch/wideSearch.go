package main

import (
	"fmt"
)

func main() {
	persons := map[string]map[string]string{
		"you": map[string]string{
			"Alice":  "noSeller",
			"Bob":    "noSeller",
			"Claire": "noSeller",
		},
		"Bob": map[string]string{
			"Anuj":  "noSeller",
			"Peggy": "noSeller",
		},
		"Alice": map[string]string{
			"Peggy": "Seller",
		},
		"Claire": map[string]string{
			"Thom":  "noSeller",
			"Jonny": "noSeller",
		},
		"Anuj":  map[string]string{},
		"Peggy": map[string]string{},
		"Thom":  map[string]string{},
		"Jonny": map[string]string{},
	}
	for k, v := range persons {
		fmt.Println(k, v)
	}
}
