package slb_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/ureuzy/acos-client-go/pkg/mocks"
	"github.com/ureuzy/acos-client-go/utils"
)

func TestGetServer(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("slb/server/myserver").Return(resp, nil)

	res, err := c.Slb.Server.Get("myserver")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetVirtualServer(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("slb/virtual-server/myvirtualserver").Return(resp, nil)

	res, err := c.Slb.VirtualServer.Get("myvirtualserver")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetVirtualServerPort(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mocks.NewMockHTTPClient(mockCtrl)
	c := mocks.GetMockClient(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("slb/virtual-server/myvirtualserver/port/myport").Return(resp, nil)

	res, err := c.Slb.VirtualServerPort.Get("myport", "myvirtualserver")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}
