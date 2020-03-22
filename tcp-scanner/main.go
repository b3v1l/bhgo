package main

import (
	"fmt"
	"net"
)

func workers(ports, results chan int) {
	for p := range ports {
		addr := fmt.Sprintf("127.0.0.1:%d", p)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	defer close(ports)
	defer close(results)

	for i := 0; i < cap(ports); i++ {

		go workers(ports, results)
	}

	go func() {
		for i := 0; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {

		port := <-results
		if port != 0 {

			openports = append(openports, port)
		}
	}
	for _, v := range openports {
		fmt.Printf("[+] Port %v open\n", v)
	}

}
