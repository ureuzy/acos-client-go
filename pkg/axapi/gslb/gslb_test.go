package gslb_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/ureuzy/acos-client-go/pkg/client"
	"github.com/ureuzy/acos-client-go/pkg/mocks"
	"github.com/ureuzy/acos-client-go/utils/testutils"
)

var cfg = client.Config{Host: "host", User: "user", Pass: "pwd", Debug: false}

func TestGetPolicy(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET("gslb/policy/mypolicy").Return(testutils.ResponseOK(), nil)

	res, err := c.Gslb.Policy.Get("mypolicy")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetServiceIP(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET("gslb/service-ip/myip").Return(testutils.ResponseOK(), nil)

	res, err := c.Gslb.ServiceIP.Get("myip")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetSite(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET("gslb/site/mysite").Return(testutils.ResponseOK(), nil)

	res, err := c.Gslb.Site.Get("mysite")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetSiteIPServer(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET("gslb/site/mysite/ip-server/myserver").Return(testutils.ResponseOK(), nil)

	res, err := c.Gslb.SiteIPServer.Get("mysite", "myserver")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetZone(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET("gslb/zone/myzone").Return(testutils.ResponseOK(), nil)

	res, err := c.Gslb.Zone.Get("myzone")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetZoneService(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET("gslb/zone/myzone/service/myport+myservice").Return(testutils.ResponseOK(), nil)

	res, err := c.Gslb.ZoneService.Get("myzone", "myport+myservice")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}
