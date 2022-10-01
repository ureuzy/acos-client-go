package rest_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/zone"
	"github.com/ureuzy/acos-client-go/pkg/client"
	"github.com/ureuzy/acos-client-go/pkg/mocks"
	"github.com/ureuzy/acos-client-go/utils/testutils"
)

var cfg = client.Config{Host: "host", User: "user", Pass: "pwd", Debug: false}

func TestCorrectArgsList(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET(gomock.Any()).Return(testutils.ResponseOK(), nil)

	_, err := c.Gslb.Zone.List()
	Ω(err).ShouldNot(HaveOccurred())
}

func TestCorrectArgsGet(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().GET(gomock.Any()).Return(testutils.ResponseOK(), nil)

	_, err := c.Gslb.Zone.Get("name")
	Ω(err).ShouldNot(HaveOccurred())
}

func TestCorrectArgsCreate(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().POST(gomock.Any(), gomock.Any()).Return(testutils.ResponseOK(), nil)

	_, err := c.Gslb.Zone.Create(&zone.Body{})
	Ω(err).ShouldNot(HaveOccurred())
}

func TestCorrectArgsCreateList(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().POST(gomock.Any(), gomock.Any()).Return(testutils.ResponseOK(), nil)

	_, err := c.Gslb.Zone.CreateList(&zone.ListBody{})
	Ω(err).ShouldNot(HaveOccurred())
}

func TestCorrectArgsDelete(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().DELETE(gomock.Any()).Return(testutils.ResponseOK(), nil)

	err := c.Gslb.Zone.Delete("name")
	Ω(err).ShouldNot(HaveOccurred())
}

func TestCorrectArgsModify(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().POST(gomock.Any(), gomock.Any()).Return(testutils.ResponseOK(), nil)

	_, err := c.Gslb.Zone.Modify(&zone.Body{}, "name")
	Ω(err).ShouldNot(HaveOccurred())
}

func TestCorrectArgsReplace(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	httpc.EXPECT().PUT(gomock.Any(), gomock.Any()).Return(testutils.ResponseOK(), nil)

	_, err := c.Gslb.Zone.Replace(&zone.Body{}, "name")
	Ω(err).ShouldNot(HaveOccurred())
}

func TestTooManyArgsGet(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	_, err := c.Gslb.ZoneService.Get("myzone", "myport+myservice", "myport+myservice")
	Ω(err).Should(HaveOccurred())
}

func TestTooFewArgsGet(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	_, err := c.Gslb.ZoneService.Get("myzone")
	Ω(err).Should(HaveOccurred())
}

func TestTooManyArgsCreateList(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	_, err := c.Gslb.Zone.CreateList(&zone.ListBody{}, "mylist")
	Ω(err).Should(HaveOccurred())
}

func TestTooFewArgsDelete(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc, cfg)

	err := c.Gslb.Zone.Delete()
	Ω(err).Should(HaveOccurred())
}
