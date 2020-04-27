package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

//3. Добавьте для time-сервера возможность его завершения при вводе команды exit.

func main() {

	// параллельный time-server
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go exit(conn)

	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

	}
}

func exit(c net.Conn) {
	cancel := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancel <- 1
	}()

	fmt.Println("Write 'exit' and hit [Enter] to exit")
	tick := make(<-chan net.Conn)
	for {
		select {
		case <-tick:
			go handleConn(c)
		case <-cancel:
			fmt.Println("Server is end")
			return
		}
	}

}
