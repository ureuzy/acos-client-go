package main

import (
	"log"

	"github.com/ureuzy/acos-client-go/pkg/axapi/slb/virtualserver"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb/virtualserver/virtualserverport"
	"github.com/ureuzy/acos-client-go/pkg/client"
)

func main() {
	config := client.Config{Host: "", User: "", Pass: "", Debug: false}
	c, err := client.NewAuthenticated(config)
	if err != nil {
		log.Fatal(err)
	}

	name := "ureuzy-sample"
	ip := "192.168.0.10"

	virtualServer, err := c.Slb.VirtualServer.Create(&virtualserver.Body{
		Object: virtualserver.Object{Name: name, IPAddress: ip},
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Slb.VirtualServerPort.CreateList(&virtualserverport.ListBody{
		ListObjects: virtualserverport.ListObjects{
			virtualserverport.Object{PortNumber: 80, Protocol: "http"},
			virtualserverport.Object{PortNumber: 443, Protocol: "https"},
		},
	}, virtualServer.Name)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Slb.VirtualServer.Delete(virtualServer.Name)
	if err != nil {
		log.Fatal(err)
	}
}
