package main

import (
	"flag"
	"os"
)

func main() {
	pName := os.Args[0]
	original := flag.String("from", "default value", "Name of original")
	copy := flag.String("copy", "default", "Name of copy")
	flag.Parse()

}
