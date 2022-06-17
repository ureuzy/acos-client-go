package slb

import (
	"fmt"

	"github.com/masanetes/acos-client-go/pkg/axapi/errors"
)

// Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server.html#virtual-server-specification

type VirtualServerListBody struct {
	VirtualServerList `json:"virtual-server-list"`
}

type VirtualServerBody struct {
	VirtualServer `json:"virtual-server"`
}

func (o *operator) ListVirtualServer() (*VirtualServerList, error) {
	res, err := o.GET("slb/virtual-server")
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response VirtualServerListBody
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServerList, nil
}

func (o *operator) GetVirtualServer(name string) (*VirtualServer, error) {
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

	var response VirtualServerBody
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServer, nil
}

func (o *operator) CreateVirtualServer(virtualServer *VirtualServerBody) (*VirtualServer, error) {
	err := errors.EmptyStringError(virtualServer.Name)
	if err != nil {
		return nil, err
	}

	res, err := o.POST("slb/virtual-server", virtualServer)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response VirtualServerBody
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServer, nil
}

func (o *operator) ModifyVirtualServer(name string, virtualServer *VirtualServerBody) (*VirtualServer, error) {
	err := errors.EmptyStringError(name)
	if err != nil {
		return nil, err
	}

	res, err := o.POST(fmt.Sprintf("slb/virtual-server/%s", name), virtualServer)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response VirtualServerBody
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServer, nil
}

func (o *operator) DeleteVirtualServer(name string) error {
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
