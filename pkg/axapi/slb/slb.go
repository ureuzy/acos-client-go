package slb

import (
	"github.com/masanetes/acos-client-go/pkg/axapi/slb/virtualserver"
	"github.com/masanetes/acos-client-go/utils"
)

type Operator struct {
	VirtualServer virtualserver.Operator
}

func New(c utils.HttpClient) *Operator {
	return &Operator{
		virtualserver.New(c),
	}
}
