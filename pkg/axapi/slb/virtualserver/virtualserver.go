package virtualserver

import (
	"fmt"

	"github.com/masanetes/acos-client-go/pkg/axapi/errors"
	"github.com/masanetes/acos-client-go/utils"
)

// Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server.html#virtual-server-specification

type operator struct {
	utils.HttpClient
}

type Operator interface {
	List() (*VirtualServerList, error)
	Get(name string) (*VirtualServer, error)
	Create(object *Object) (*VirtualServer, error)
	Modify(name string, object *Object) (*VirtualServer, error)
	Delete(name string) error
}

func New(c utils.HttpClient) Operator {
	return &operator{c}
}

func (o *operator) List() (*VirtualServerList, error) {
	res, err := o.GET("slb/virtual-server")
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response ListObject
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServerList, nil
}

func (o *operator) Get(name string) (*VirtualServer, error) {
	err := errors.EmptyStringError(name)
	if err != nil {
		return nil, err
	}

	res, err := o.GET(fmt.Sprintf("slb/virtual-server/%s", name))
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response Object
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServer, nil
}

func (o *operator) Create(object *Object) (*VirtualServer, error) {
	err := errors.EmptyStringError(object.Name)
	if err != nil {
		return nil, err
	}

	res, err := o.POST("slb/virtual-server", object)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response Object
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServer, nil
}

func (o *operator) Modify(name string, object *Object) (*VirtualServer, error) {
	err := errors.EmptyStringError(name)
	if err != nil {
		return nil, err
	}

	res, err := o.POST(fmt.Sprintf("slb/virtual-server/%s", name), object)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response Object
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServer, nil
}

func (o *operator) Delete(name string) error {
	err := errors.EmptyStringError(name)
	if err != nil {
		return err
	}

	res, err := o.DELETE(fmt.Sprintf("slb/virtual-server/%s", name))
	if err != nil {
		return err
	}

	if res.HasError() {
		return errors.Handle(res)
	}

	return nil
}
