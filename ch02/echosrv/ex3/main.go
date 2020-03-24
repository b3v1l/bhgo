package main

import (
	"bufio"
	"log"
	"net"
)

// same echo server but using bufio

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Error in reading data")
	}
	log.Printf("Received %d bytes: %s \n", len(s), s)

	_, err = writer.WriteString(s)
	if err != nil {
		log.Fatalln("Cannot write data")
	}
	writer.Flush()
}

func main() {

	//bind a connection
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalln("Error setting up the Listener")
	}
	log.Println("Listening on port 8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")

		}
		log.Println("Received connection from %v", conn)
		go echo(conn)
	}

}
