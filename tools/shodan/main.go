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

	//	query, err := s.HostSearch(os.Args[1])
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(query)
	//	fmt.Println(query.Matches)
	//
	//	fmt.Println("test!")
	//	for _, host := range query.Matches {
	//		fmt.Printf("IP=%s", host.IPString) // host.Port)
	//	}
	//
}

//
