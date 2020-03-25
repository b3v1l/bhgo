package main

import (
	"bhgo/tools/shodan/shodan"
	"fmt"
	"os"
)

func main() {

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(
		"[+] Query Credits left = %d\n[+] Scan_Credits left = %d\n", info.Querycredits, info.Scancredits)
}
