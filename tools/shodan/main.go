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
	// Organisation search

	//flag.Usage()

	orgCmd := flag.NewFlagSet("org", flag.ContinueOnError)
	orgGet := orgCmd.String("h", "", "Usage: -h ORG. Retrieved information about target Organisation")

	// Domain name
	dnsreq := flag.NewFlagSet("dns", flag.ExitOnError)
	dns := dnsreq.String("d", "", "Usage: -dns 'Domain Name'")

	domCmd := flag.NewFlagSet("domain", flag.ContinueOnError)
	domGet := domCmd.String("d", "", "Usage: domain -d 'Domain name'")

	//Dns Reverse resolution

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		fmt.Println("Usage: shodan-scan <command> [<args>]")
		fmt.Println("Commands are 'ip - dns - domain'")
		os.Exit(2)
		//flag.Usage()
		//log.Fatalln("Usage: shodan Searchterm")
		//		os.Exit(1)
	}

	ipCmd := flag.NewFlagSet("ip", flag.ExitOnError)
	ipGet := ipCmd.String("h", "", "Usage: ip -h 1.2.3.4")

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
		OrgSearch(*orgGet)
		fmt.Println(len(os.Args[2:]))
		if len(os.Args[2:]) == 0 {
			//fmt.Println("Usage: org -h organisation")
			os.Exit(1)
		}

	case "dns":
		dnsreq.Parse(os.Args[2:])
		Dns(*dns)
		if len(os.Args[2:]) <= 1 {
			//fmt.Println("Usage: org -h organisation")
			os.Exit(1)
		}

	case "ip":
		ipCmd.Parse(os.Args[2:])
		query, err := s.HostIP(*ipGet)
		if err != nil {
			log.Fatalln(err)
		}
		//for _, v := range query.Data {
		//	fmt.Printf("Ip %v", v)
		//}
		fmt.Printf("Hosting %v\n", query.Data[0:])

	case "domain":

		domCmd.Parse(os.Args[2:])
		Domain(*domGet)

	default:
		fmt.Println("Usage: shodan-scan <command> [<args>]")
		fmt.Println("Commands are 'ip - dns - domain'")
		os.Exit(2)
	}

}

func OrgSearch(org string) {

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	query, err := s.HostSearch(org)
	if err != nil {
		fmt.Println(err)
	}
	if len(os.Args[2:]) == 0 {
		os.Exit(1)
	} else {
		fmt.Printf("[*] Results for target '%v'\n\n", org)
		for _, host := range query.Matches {
			//		index := len(host.Domains)
			for _, hosting := range host.Domains {
				fmt.Printf("IP:%18s Port:%8d Hosting:%30s Location: %v\n", host.IPString, host.Port, hosting, host.Location)
			}
		}
	}
}
func Dns(dns string) {

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	hostlist := make([]string, 0)
	hostlist = append(hostlist, dns)
	dsearch, err := s.DnsSearch(hostlist)
	time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range dsearch {
		fmt.Printf("DNS Resolution  : %v\n", v)
	}
}

func Domain(dom string) {

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)

	resp, err := s.DomainInfo(dom)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Result for domain %v\n", resp.Domain)
	fmt.Printf("Result for domain %v\n", resp.Tags)

	for _, v := range resp.Data {
		fmt.Printf("%s\n", v)
	}
}
