package health

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/health/healthmonitor"
	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

const path = "health"

type Operator struct {
	Montitor rest.Operator[healthmonitor.Body, healthmonitor.ListBody]
}

func New(c utils.HTTPClient) *Operator {
	return &Operator{
		Montitor: healthmonitor.New(c, path),
	}
}
