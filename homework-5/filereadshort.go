package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	sh, err := ioutil.ReadFile("filereadshort.go")
	if err != nil {
		return
	}
	fmt.Println(string(sh))
}
