package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	req, err := http.Get("https://google.be/robots.txt")
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(req.Status)

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(b))
	req.Body.Close()

}
