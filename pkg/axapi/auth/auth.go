package auth

import (
	"github.com/masanetes/acos-client-go/utils"
)

type Request struct {
	Credentials `json:"credentials"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	*AuthResponse `json:"authresponse"`
}

type AuthResponse struct {
	Signature   string `json:"signature"`
	Description string `json:"description"`
}

type operator struct {
	utils.HttpClient
}

type Operator interface {
	Request(req *Request) (*AuthResponse, error)
}

func New(c utils.HttpClient) Operator {
	return &operator{c}
}

func (o *operator) Request(req *Request) (*AuthResponse, error) {
	res, err := o.POST("/auth", req)
	if err != nil {
		return nil, err
	}

	var response Response
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return response.AuthResponse, nil
}
