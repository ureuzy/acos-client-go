package mocks

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/activepartition"
	"github.com/ureuzy/acos-client-go/pkg/axapi/auth"
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb"
	"github.com/ureuzy/acos-client-go/pkg/axapi/health"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb"
	"github.com/ureuzy/acos-client-go/pkg/client"
)

func GetMockClient(httpc *MockHTTPClient) *client.Client {
	return &client.Client{
		Auth:            auth.New(httpc),
		ActivePartition: activepartition.New(httpc),
		Health:          health.New(httpc),
		Slb:             slb.New(httpc),
		Gslb:            gslb.New(httpc),
	}
}
