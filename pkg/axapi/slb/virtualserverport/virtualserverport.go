package virtualserverport

import (
	"fmt"

	"github.com/masanetes/acos-client-go/pkg/axapi/errors"
	"github.com/masanetes/acos-client-go/utils"
)

// Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#port-specification

type operator struct {
	utils.HttpClient
}

type Operator interface {
	List(virtualServerName string) (*VirtualServerPortList, error)
	Get(virtualServerName string, protocol string, portNumber int) (*VirtualServerPort, error)
	Create(virtualServerName string, object *Object) (*VirtualServerPort, error)
	CreateList(virtualServerName string, object *ListObject) (*VirtualServerPortList, error)
	Modify(virtualServerName string, protocol string, portNumber int, object *Object) (*VirtualServerPort, error)
	Delete(virtualServerName string, protocol string, portNumber int) error
}

func New(c utils.HttpClient) Operator {
	return &operator{c}
}

func (o *operator) List(virtualServerName string) (*VirtualServerPortList, error) {
	res, err := o.GET(fmt.Sprintf("slb/virtual-server/%s/port", virtualServerName))
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

	return &response.VirtualServerPortList, nil
}

func (o *operator) Get(virtualServerName string, protocol string, portNumber int) (*VirtualServerPort, error) {
	err := errors.EmptyStringError(virtualServerName)
	if err != nil {
		return nil, err
	}

	res, err := o.GET(fmt.Sprintf("slb/virtual-server/%s/port/%d+%s",
		virtualServerName,
		portNumber,
		protocol,
	))
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

	return &response.VirtualServerPort, nil
}

func (o *operator) Create(virtualServerName string, object *Object) (*VirtualServerPort, error) {
	err := errors.EmptyStringError(virtualServerName)
	if err != nil {
		return nil, err
	}

	res, err := o.POST(fmt.Sprintf("slb/virtual-server/%s/port", virtualServerName), object)
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

	return &response.VirtualServerPort, nil
}

func (o *operator) CreateList(virtualServerName string, object *ListObject) (*VirtualServerPortList, error) {
	err := errors.EmptyStringError(virtualServerName)
	if err != nil {
		return nil, err
	}

	res, err := o.POST(fmt.Sprintf("slb/virtual-server/%s/port", virtualServerName), object)
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

	return &response.VirtualServerPortList, nil
}

func (o *operator) Modify(virtualServerName string, protocol string, portNumber int, object *Object) (*VirtualServerPort, error) {
	err := errors.EmptyStringError(virtualServerName)
	if err != nil {
		return nil, err
	}

	res, err := o.POST(fmt.Sprintf("slb/virtual-server/%s/port/%d+%s",
		virtualServerName,
		portNumber,
		protocol,
	), object)
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

	return &response.VirtualServerPort, nil
}

func (o *operator) Delete(virtualServerName string, protocol string, portNumber int) error {
	err := errors.EmptyStringError(virtualServerName)
	if err != nil {
		return err
	}

	res, err := o.DELETE(fmt.Sprintf("slb/virtual-server/%s/port/%d+%s",
		virtualServerName,
		portNumber,
		protocol,
	))
	if err != nil {
		return err
	}

	if res.HasError() {
		return errors.Handle(res)
	}

	return nil
}
