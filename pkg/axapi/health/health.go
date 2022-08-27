package health

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/health/monitor"
	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

const path = "health"

type Operator struct {
	Montitor rest.Operator[monitor.Body, monitor.ListBody]
}

func New(c utils.HTTPClient) *Operator {
	return &Operator{
		Montitor: monitor.New(c, path),
	}
}
