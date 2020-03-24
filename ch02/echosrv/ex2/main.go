package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 512)
	for {
		size, err := conn.Read(buffer[0:])
		if err == io.EOF {
			log.Fatalln("Client disconnected")
			break
		}
		if err != nil {
			log.Fatalln("Error, closing connection")
			break
		}
		log.Printf("Received %d from output %s", size, string(buffer))

		//Send data using Write
		log.Println("Writing data")

		if _, err := conn.Write(buffer[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}

	}
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
		log.Println("Received connection from %v", &conn)
		go echo(conn)
	}

}
