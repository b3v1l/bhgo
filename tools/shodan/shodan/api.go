package shodan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIInfo struct {
	Scancredits  int    `json:"scan_credits"`
	Plan         string `json:"plani"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
	Querycredits int    `json:"query_credits"`
	Monitoredips string `json:monitored_ips`
	Unlockedleft int    `json:"unlocked_left"`
	Telnet       bool   `json:"telnet"`
}

//func get aip info

func (c *Client) APIInfo() (*APIInfo, error) {

	resp, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, c.apiKey))
	if err != nil {
		log.Fatalln(err)
	}
	//	fmt.Println(resp)

	var ret APIInfo
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &ret, err

}
