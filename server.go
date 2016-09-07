package main // import "github.com/deviantony/gosrv"

import (
	"log"
	"net"
)

const (
	port = "7777"
)

var clients []net.Conn

func main() {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Print("Server listening...")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
		}

		log.Printf("Client connected %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Add the client to the connection array
		clients = append(clients, conn)

		go handler(conn)
	}
}

func removeClient(conn net.Conn) {
	log.Printf("Client %s disconnected", conn.RemoteAddr())
	conn.Close()
	//remove client from clients here
}

func handler(conn net.Conn) {
	defer removeClient(conn)
	errorChan := make(chan error)
	dataChan := make(chan []byte)

	go readWrapper(conn, dataChan, errorChan)

	for {
		select {
		case data := <-dataChan:
			log.Printf("Client %s sent: %s", conn.RemoteAddr(), string(data))
			for i := range clients {
				clients[i].Write(data)
			}
		case err := <-errorChan:
			log.Println("An error occured:", err.Error())
			return
		}
	}
}

func readWrapper(conn net.Conn, dataChan chan []byte, errorChan chan error) {
	for {
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			errorChan <- err
			return
		}
		dataChan <- buf[:reqLen]
	}
}
