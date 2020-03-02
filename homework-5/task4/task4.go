package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var pName = os.Args[0]
var original = flag.String("from", "", "Name of original")
var copy = flag.String("copy", "", "Name of copy")

//help for user
func help() {
	fmt.Println("чтобы скопировать файл введите в консоль")
	fmt.Printf("%s -original=original-name.txt -copy=copy-name.txt\n", pName)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	//если один из параметров не указан, то вывыедет справку
	if *original == "" || *copy == "" {
		help()
	}

	// open whole file, read and close after
	//прочитать весь файл будет лучше, т.к. мы хотим скопировать весь файл целиком
	//если же будем читать построчно или побайтово, то мало ли что-то недокопируется
	// не знаю такое может быть или нет,но пока это вижу так
	fData, err := ioutil.ReadFile(*original)
	if err != nil {
		log.Fatal(err)
	}

	nFile, err := os.Create(*copy)
	if err != nil {
		log.Fatal(err)
	}
	defer nFile.Close()
	nFile.WriteString(string(fData))

}
