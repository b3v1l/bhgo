package main

import (
	"bhgo/tools/shodan/shodan"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan Searchterm")
	}

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(
		"[+] Query Credits left = %d\n[+] Scan_Credits left = %d\n",
		info.Querycredits, info.Scancredits)

	query, err := s.HostSearch(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[*] Results for target '%v'\n", os.Args[1])
	for _, host := range query.Matches {
		fmt.Printf("IP:%18s Port:%8d\n", host.IPString, host.Port)
	}
}
