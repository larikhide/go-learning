package main

import (
	"fmt"
	"io"
	"os"
)

func dirTree(decrOfFile io.Writer, nameOfSomth string, hzChto bool) error {
	dir, err := os.Open(".")
	if err != nil {
		return fmt.Errorf("file not sorted")
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("end of directory")
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	return err
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
