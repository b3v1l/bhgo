package main

import (
	"bhgo/tools/shodan/shodan"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	orgCmd := flag.NewFlagSet("org", flag.ExitOnError)
	orgGet := orgCmd.String("h", "", "Usage: -h ORG. Retrieved information about target Organisation")

	dnsreq := flag.NewFlagSet("dns", flag.ExitOnError)
	dns := dnsreq.String("d", "", "Usage: -dns 'Domain Name'")

	if len(os.Args) < 2 {
		log.Fatalln("Usage: shodan Searchterm")
		os.Exit(1)
	}

	//fmt.Println(*dns)
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(
		"[+] Query Credits left = %d\n[+] Scan_Credits left = %d\n",
		info.Querycredits, info.Scancredits)

	//flag.Parse()

	switch os.Args[1] {
	case "org":

		orgCmd.Parse(os.Args[2:])
		query, err := s.HostSearch(*orgGet)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("[*] Results for target '%v'\n", *orgGet)
		for _, host := range query.Matches {
			//		index := len(host.Domains)
			for _, hosting := range host.Domains {
				fmt.Printf("IP:%18s Port:%8d Domain:%40s\n", host.IPString, host.Port, hosting)
			}
		}

	case "dns":
		dnsreq.Parse(os.Args[2:])
		hostlist := make([]string, 0)
		hostlist = append(hostlist, *dns)
		dsearch, err := s.DnsSearch(hostlist)
		time.Sleep(2 * time.Second)
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range dsearch {
			fmt.Printf("DNS Resolution  : %v\n", v)
		}

	default:
		fmt.Println("Error")
	}

}
