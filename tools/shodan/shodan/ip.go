package shodan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HostIP struct {
	RegionCode   string      `json:"region_code"`
	IP           int         `json:"ip"`
	AreaCode     int         `json:"area_code"`
	Latitude     float64     `json:"latitude"`
	Hostnames    []string    `json:"hostnames"`
	PostalCode   string      `json:"postal_code"`
	DmaCode      int         `json:"dma_code"`
	CountryCode  string      `json:"country_code"`
	Org          string      `json:"org"`
	Data         []Data      `json:"data"`
	City         string      `json:"city"`
	Isp          string      `json:"isp"`
	Longitude    float64     `json:"longitude"`
	LastUpdate   string      `json:"last_update"`
	CountryCode3 string      `json:"country_code3"`
	CountryName  string      `json:"country_name"`
	IPStr        string      `json:"ip_str"`
	Os           interface{} `json:"os"`
	Ports        []int       `json:"ports"`
}

type Data struct {
	//Product string `json:"product"`
	//Title   string `json:"title"`

	//Timestamp string   `json:"timestamp"`
	//Isp  string   `json:"isp"`
	//Cpe  []string `json:"cpe"`
	//Data string   `json:"data"`
	//	HTML string   `json:"html"`

	IP        int         `json:"ip"`
	Domains   []string    `json:"domains"`
	Org       string      `json:"org"`
	Os        interface{} `json:"os"`
	Port      int         `json:"port"`
	Hostnames []string    `json:"hostnames"`
	IPStr     string      `json:"ip_str"`
}

func (c *Client) HostIP(ip string) (*HostIP, error) {

	h := &HostIP{}
	//https://api.shodan.io/shodan/host/{ip}?key={YOUR_API_KEY}
	resp, err := http.Get(fmt.Sprintf("%s/shodan/host/%s?key=%s", BaseURL, ip, c.apiKey))
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&h); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return h, nil
}
