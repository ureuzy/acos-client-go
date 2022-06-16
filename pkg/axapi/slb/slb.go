package slb

import (
	"github.com/masanetes/acos-client-go/utils"
)

type operator struct {
	utils.HttpClient
}

type Operator interface {
	ListVirtualServer() (*VirtualServerList, error)
	GetVirtualServer(name string) (*VirtualServer, error)
}

func New(c utils.HttpClient) Operator {
	return &operator{c}
}
