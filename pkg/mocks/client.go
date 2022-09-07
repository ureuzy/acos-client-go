package mocks

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/activepartition"
	"github.com/ureuzy/acos-client-go/pkg/axapi/auth"
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb"
	"github.com/ureuzy/acos-client-go/pkg/axapi/health"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb"
	"github.com/ureuzy/acos-client-go/pkg/client"
)

func GetMockClient(httpc *MockHTTPClient, cfg client.Config) *client.Client {
	c := client.Instance(cfg)
	c.Auth = auth.New(httpc)
	c.ActivePartition = activepartition.New(httpc)
	c.Health = health.New(httpc)
	c.Slb = slb.New(httpc)
	c.Gslb = gslb.New(httpc)
	return c
}
