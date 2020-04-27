package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//net.Dial создает подключение к серверу в указанной сети tcp
	//по адресу localhost:8000
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
// после завершения главной горутины закрыть соединение
	defer conn.Close()

	// копирование в стандартный вывод
	go func() {
		io.Copy(os.Stdout, conn)
	}()


	// копирование из ввода в conn
	io.Copy(conn, os.Stdin) // until you send ^Z
	fmt.Printf("%s: exit", conn.LocalAddr())
}
