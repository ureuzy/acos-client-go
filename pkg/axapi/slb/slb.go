package slb

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb/server"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb/virtualserver"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb/virtualserverport"
	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

const path = "slb"

type Operator struct {
	Server            rest.Operator[server.Body, server.ListBody]
	VirtualServer     rest.Operator[virtualserver.Body, virtualserver.ListBody]
	VirtualServerPort rest.Operator[virtualserverport.Body, virtualserverport.ListBody]
}

func New(c utils.HTTPClient) *Operator {
	return &Operator{
		Server:            server.New(c, path),
		VirtualServer:     virtualserver.New(c, path),
		VirtualServerPort: virtualserverport.New(c, path),
	}
}
