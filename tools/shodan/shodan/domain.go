package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DomainInfo struct {
	Domain string    `json: "domain"`
	Tags   []string  `json: "tags"`
	Data   []SubData `json: "data"`
}

type SubData struct {
	Subdomain string `json: "subdomain"`
	Type      string `json: "type"`
	Value     string `json: "value"`
	LastSeen  string `json: "last_seen"`
}

func (c *Client) DomainInfo(query string) (*DomainInfo, error) {

	resp, err := http.Get(fmt.Sprintf("%s/dns/domain/%s?key=%s", BaseURL, query, c.apiKey))
	if err != nil {
		return nil, err
	}

	var ret DomainInfo

	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {

		return nil, err
	}
	defer resp.Body.Close()

	return &ret, nil
}
