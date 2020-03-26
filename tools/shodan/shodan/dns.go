package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DnsSearch struct {
	Hostname string
	IP       string
}

func (c *Client) DnsSearch(hostnames []string) ([]DnsSearch, error) {

	//https://api.shodan.io/dns/resolve?hostnames=fedasil&key=\
	//tab := make(map[string]string)
	d := []DnsSearch{}

	resp, err := http.Get(fmt.Sprintf("%s/dns/resolve?hostnames=%s&key=%s", BaseURL, strings.Join(hostnames, ","), c.apiKey))
	if err != nil {
		return d, err
	}
	defer resp.Body.Close()
	m := make(map[string]string)

	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, err
	}

	for k, v := range m {
		d = append(d, DnsSearch{
			Hostname: k,
			IP:       v,
		})
	}
	return d, nil
}
