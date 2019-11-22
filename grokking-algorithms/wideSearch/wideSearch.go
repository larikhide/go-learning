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
			"Thom":  "Seller",
			"Jonny": "noSeller",
		},
		"Anuj":  map[string]string{},
		"Peggy": map[string]string{},
		"Thom":  map[string]string{},
		"Jonny": map[string]string{},
	}
	for _, v := range persons {
		for name, vs := range v {
			if vs == "Seller" {
				fmt.Printf("%s is %s of mango\n", name, vs)
			}
		}
	}
}
