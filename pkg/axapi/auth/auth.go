package auth

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/errors"
	"github.com/ureuzy/acos-client-go/utils"
)

type operator struct {
	utils.HTTPClient
}

type Operator interface {
	Login(req *Request) (*Body, error)
	Logoff() error
}

func New(c utils.HTTPClient) Operator {
	return &operator{HTTPClient: c}
}

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

func (o *operator) Logoff() error {
	_, err := o.POST("logoff", "{}")
	return err
}
