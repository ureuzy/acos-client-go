package slb

import (
	"fmt"
	"github.com/masanetes/acos-client-go/utils"
)

type VirtualServerListResponse struct {
	VirtualServerList `json:"virtual-server-list"`
}

type VirtualServerResponse struct {
	VirtualServer `json:"virtual-server"`
}

func (o *operator) ListVirtualServer() (*VirtualServerList, error) {
	res, err := o.GET("slb/virtual-server")
	if err != nil {
		return nil, err
	}

	var response VirtualServerListResponse
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServerList, nil
}

func (o *operator) GetVirtualServer(name string) (*VirtualServer, error) {

	err := utils.EmptyStringError(name)
	if err != nil {
		return nil, err
	}

	res, err := o.GET(fmt.Sprintf("slb/virtual-server/%s", name))
	if err != nil {
		return nil, err
	}

	var response VirtualServerResponse
	if err = res.UnmarshalJson(&response); err != nil {
		return nil, err
	}

	return &response.VirtualServer, nil
}
