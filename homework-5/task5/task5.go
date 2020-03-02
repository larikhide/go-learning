package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var pName = os.Args[0]
var file = flag.String("f", "", "Файл, в котором ищем")

func help() {
	fmt.Printf("  %s -f='file.ext' 'search text'\n", pName)
}

func findText(path string, args string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//чтение построчно
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		if strings.Contains(line, args) {
			fmt.Printf("%s - обнаружено на %d строке", line, i)
		}
	}

}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 || args[0] == "" {
		help()
	}

	findText(*file, args[0])

}
