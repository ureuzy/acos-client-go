package auth

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/errors"
	"github.com/ureuzy/acos-client-go/utils"
)

type operator struct {
	utils.HTTPClient
}

// Operator is an operation related to authentication
type Operator interface {
	Login(req *Request) (*Body, error)
	Logoff() error
}

func New(c utils.HTTPClient) Operator {
	return &operator{HTTPClient: c}
}

// Login Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/start_here.html#logging-on
func (o *operator) Login(req *Request) (*Body, error) {
	res, err := o.POST("auth", req)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response Response
	if err = res.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return response.Body, nil
}

// Logoff Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/start_here.html#logging-off
func (o *operator) Logoff() error {
	_, err := o.POST("logoff", "{}")
	return err
}
