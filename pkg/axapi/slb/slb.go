package slb

import (
	"github.com/masanetes/acos-client-go/pkg/axapi/slb/server"
	"github.com/masanetes/acos-client-go/pkg/axapi/slb/virtualserver"
	"github.com/masanetes/acos-client-go/pkg/axapi/slb/virtualserverport"
	"github.com/masanetes/acos-client-go/utils"
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
