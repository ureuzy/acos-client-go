package main

import (
	"log"

	"github.com/masanetes/acos-client-go/pkg/client"
)

func main() {

	config := client.Config{Host: "192.168.0.10", User: "sample", Pass: "sample", Debug: false}
	c, err := client.New(&config, client.InsecureSkipVerify(true))
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Slb.GetVirtualServer("masanetes-sample-virtualserver")
	if err != nil {
		log.Fatal(err)
	}
}
