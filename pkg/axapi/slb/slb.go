package slb

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb/server"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb/virtualserver"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb/virtualserverport"
	"github.com/ureuzy/acos-client-go/utils"
)

type Operator struct {
	Server            server.Operator
	VirtualServer     virtualserver.Operator
	VirtualServerPort virtualserverport.Operator
}

func New(c utils.HttpClient) *Operator {
	return &Operator{
		server.New(c),
		virtualserver.New(c),
		virtualserverport.New(c),
	}
}
