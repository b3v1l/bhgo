package main

import (
	"bhgo/tools/shodan/shodan"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

//Global key API variable
var apiKey = os.Getenv("SHODAN_API_KEY")
var s = shodan.New(apiKey)

func main() {
	// Organisation search

	//flag.Usage()

	orgCmd := flag.NewFlagSet("org", flag.ExitOnError)
	orgGet := orgCmd.String("h", "", "Usage: -h ORG. Retrieved information about target Organisation")

	// Domain name
	dnsreq := flag.NewFlagSet("dns", flag.ExitOnError)
	dns := dnsreq.String("d", "", "Usage: -dns 'Domain Name'")

	domCmd := flag.NewFlagSet("domain", flag.ExitOnError)
	domGet := domCmd.String("d", "", "Usage: domain -d 'Domain name'")

	ipCmd := flag.NewFlagSet("ip", flag.ExitOnError)
	ipGet := ipCmd.String("h", "", "Usage: ip -h 1.2.3.4")
	//Dns Reverse resolution

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		fmt.Println("Usage: shodan-scan <command> [<args>]")
		fmt.Println("	Commands are 'org - ip - dns - domain'")
		os.Exit(2)

	}

	switch os.Args[1] {
	case "org":

		orgCmd.Parse(os.Args[2:])
		if orgCmd.Parsed() {
			// Required Flags
			if *orgGet == "" {
				orgCmd.PrintDefaults()
				os.Exit(1)
			}
		}
		OrgSearch(*orgGet)

	case "dns":
		dnsreq.Parse(os.Args[2:])
		if dnsreq.Parsed() {
			// Required Flags
			if *dns == "" {
				dnsreq.PrintDefaults()
				os.Exit(1)
			}
		}
		Dns(*dns)

	case "ip":
		ApiStat()
		ipCmd.Parse(os.Args[2:])
		if ipCmd.Parsed() {
			// Required Flags
			if *ipGet == "" {
				ipCmd.PrintDefaults()
				os.Exit(1)
			}
		}
		IpQuery(*ipGet)

	case "domain":
		ApiStat()
		domCmd.Parse(os.Args[2:])
		if domCmd.Parsed() {
			// Required Flags
			if *domGet == "" {
				domCmd.PrintDefaults()
				os.Exit(1)
			}
		}
		Domain(*domGet)

	default:
		fmt.Println("Usage: shodan-scan <command> [<args>]")
		fmt.Println("Commands are 'org - ip - dns - domain'")
		os.Exit(2)
	}
}

func ApiStat() {

	info, err := s.APIInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n**************STATS******************")
	fmt.Printf(
		"[+] Query Credits left = %d\n[+] Scan_Credits left = %d\n",
		info.Querycredits, info.Scancredits)
	fmt.Println("*************************************\n")
}

func OrgSearch(org string) {

	query, err := s.HostSearch(org)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[*] Results for target '%v'\n\n", org)
	for _, host := range query.Matches {
		//		index := len(host.Domains)
		for _, hosting := range host.Domains {
			fmt.Printf("IP:%18s Port:%8d Hosting:%30s Location: %v\n", host.IPString, host.Port, hosting, host.Location)
		}
	}

}
func Dns(dns string) {

	hostlist := make([]string, 0)
	hostlist = append(hostlist, dns)
	dsearch, err := s.DnsSearch(hostlist)
	time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range dsearch {
		fmt.Printf("DNS Resolution  : \n%v\n", v)
	}
}

func Domain(dom string) {

	resp, err := s.DomainInfo(dom)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Result for domain %v\n", resp.Domain)
	fmt.Printf("Tags %v\n", resp.Tags)

	for _, v := range resp.Data {
		fmt.Printf("%s\n", v)
	}
}

func IpQuery(q string) {

	query, err := s.HostIP(q)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range query.Data {
		fmt.Printf("Ip %v\n", v)
	}
	//fmt.Printf("Hosting %v\n", query.Data[0:])
}
