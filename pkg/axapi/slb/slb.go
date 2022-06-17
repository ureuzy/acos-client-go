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
	CreateVirtualServer(virtualServer *VirtualServerBody) (*VirtualServer, error)
	ModifyVirtualServer(virtualServerName string, virtualServer *VirtualServerBody) (*VirtualServer, error)
	DeleteVirtualServer(name string) error
}

func New(c utils.HttpClient) Operator {
	return &operator{c}
}
