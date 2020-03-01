package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)
/*
2. Что бы вы изменили в программе чтения из файла? Приведите исправленный вариант, обоснуйте свое решение в комментарии. */

// файл может быть большого размера и лучше читать его построчно или по определенному количеству байт

//чтение побайтово
func main() {

	//открыли файл
	file, err := os.Open("fileread.go")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()


	//прочли по 64 байта
	bs := make([]byte, 64)
	for{
        n, err := file.Read(bs)
        if err == io.EOF{   // если конец файла
            break           // выходим из цикла
        }
        fmt.Print(string(bs[:n]))
    }


//построчное чтение из файла
	anFile, err := os.Open("task1.txt")
    if err != nil {
         fmt.Println("Unable to open file:", err)
         return
    }
    defer anFile.Close()

    reader := bufio.NewReader(anFile)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            } else {
                fmt.Println(err)
                return
            }
        }
        fmt.Print(line)
    }
}
