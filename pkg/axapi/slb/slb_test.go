package slb_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/ureuzy/acos-client-go/pkg/client"
	"github.com/ureuzy/acos-client-go/pkg/mocks"
	"github.com/ureuzy/acos-client-go/utils/testutils"
)

var cfg = client.Config{Host: "host", User: "user", Pass: "pwd", Debug: false}

func TestGetServer(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET("slb/server/myserver").Return(testutils.ResponseOK(), nil)

	res, err := c.Slb.Server.Get("myserver")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetVirtualServer(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET("slb/virtual-server/myvirtualserver").Return(testutils.ResponseOK(), nil)

	res, err := c.Slb.VirtualServer.Get("myvirtualserver")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetVirtualServerPort(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET("slb/virtual-server/myvirtualserver/port/myport").Return(testutils.ResponseOK(), nil)

	res, err := c.Slb.VirtualServerPort.Get("myvirtualserver", "myport")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}
