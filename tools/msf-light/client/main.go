package main

import (
	"fmt"
	"log"
	"os"

	"bhgo/tools/msf-light/rpc"
)

func main() {

	host := os.Getenv("MSFHOST")
	user := "msf"
	pass := os.Getenv("MSFPASS")

	if host == "" || pass == "" {
		log.Fatalln("Error parsing username or password ...")
	}
	msf, err := rpc.New(user, pass, host)
	if err != nil {
		log.Fatalln("Error retrieving session ...")
	}
	defer msf.Logout()

	sessions, err := msf.SessionList()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Sessions:")
	for _, session := range sessions {

		fmt.Printf("%5d %s\n", session.ID, session.Info)

	}

}
