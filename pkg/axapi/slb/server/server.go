package server

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/axapi/errors"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#server-specification

type operator struct {
	utils.HttpClient
}

type Operator interface {
	List() (*ListObjects, error)
	Get(name string) (*Object, error)
	Create(object *Body) (*Object, error)
	Modify(name string, object *Body) (*Object, error)
	Delete(name string) error
}

func New(c utils.HttpClient) Operator {
	return &operator{c}
}

func (o *operator) List() (*ListObjects, error) {
	res, err := o.GET("slb/server")
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response ListBody
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.ListObjects, nil
}

func (o *operator) Get(name string) (*Object, error) {
	err := errors.EmptyStringError(name)
	if err != nil {
		return nil, err
	}

	res, err := o.GET(fmt.Sprintf("slb/server/%s", name))
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response Body
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.Object, nil
}

func (o *operator) Create(object *Body) (*Object, error) {
	err := errors.EmptyStringError(object.Name)
	if err != nil {
		return nil, err
	}

	res, err := o.POST("slb/server", object)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response Body
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.Object, nil
}

func (o *operator) Modify(name string, object *Body) (*Object, error) {
	err := errors.EmptyStringError(name)
	if err != nil {
		return nil, err
	}

	res, err := o.POST(fmt.Sprintf("slb/server/%s", name), object)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response Body
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.Object, nil
}

func (o *operator) Delete(name string) error {
	err := errors.EmptyStringError(name)
	if err != nil {
		return err
	}

	res, err := o.DELETE(fmt.Sprintf("slb/server/%s", name))
	if err != nil {
		return err
	}

	if res.HasError() {
		return errors.Handle(res)
	}

	return nil
}
