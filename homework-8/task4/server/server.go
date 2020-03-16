package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//поток который принимает
type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)


//broadcaster хранит информацию о всех клиентах и прослушивает каналы событий и сообщений, используя мультиплексирование с помощью select.
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}



/* создает новый канал исходящих сообщений для своего клиента и объявляет широковещателю о поступлении этого клиента по каналу entуring. Затем она считывает каждую
строку текста от клиента, отправляет их широковещателю по глобальному каналу входящих сообщений, предваряя каждое сообщение указанием отправителя. Когда от клиента получена вся информация, handleConn объявляет об убытии клиента по каналу leaving и закрывает подключение. */
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}


func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}


	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
