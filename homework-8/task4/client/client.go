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

	//не понял как работает данная горутина.
	//мы из conn (интерфейс) копируем что? в стандартный вывод. не понимаю что копируем
	go func() {
		io.Copy(os.Stdout, conn)
	}()

	// тот же вопрос. что именно мы посылаем на сервер и как вообще
	// происходит передача каких-то данных. мы же просто занимаемся копированием.
	// не понимаю почему закрывается соединение, у нас же даже условия никакого нет
	io.Copy(conn, os.Stdin) // until you send ^Z
	fmt.Printf("%s: exit", conn.LocalAddr())
}
