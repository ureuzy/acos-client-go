package auth

import (
	"github.com/ureuzy/acos-client-go/utils"
)

type operator struct {
	utils.HTTPClient
	basePath string
}

type Operator interface {
	Request(req *Request) (*Body, error)
}

func New(c utils.HTTPClient) Operator {
	const path = "auth"
	return &operator{HTTPClient: c, basePath: path}
}

func (o *operator) Request(req *Request) (*Body, error) {
	res, err := o.POST(o.basePath, req)
	if err != nil {
		return nil, err
	}

	var response Response
	if err = res.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return response.Body, nil
}
