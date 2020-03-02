package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	// open whole file, read and close after
	sh, err := ioutil.ReadFile("filereadshort.go")
	if err != nil {
		return
	}
	fmt.Println(string(sh))

	//read directory
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	//create a file and write something into
	file, err := os.Create("task1.txt")
	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("тут я напишу что-нибудь по первому заданию ДЗ номер 5")

	//Рекурсивный обход каталогов
	filepath.Walk("../..", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
