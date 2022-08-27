package slb_test

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	mockutils "github.com/ureuzy/acos-client-go/pkg/mocks/utils"
	"github.com/ureuzy/acos-client-go/utils"
)

func TestGetServer(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mockutils.NewMockHTTPClient(mockCtrl)
	sut := slb.New(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("slb/server/myserver").Return(resp, nil)

	res, err := sut.Server.Get("myserver")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetVirtualServer(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mockutils.NewMockHTTPClient(mockCtrl)
	sut := slb.New(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("slb/virtual-server/myvirtualserver").Return(resp, nil)

	res, err := sut.VirtualServer.Get("myvirtualserver")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetVirtualServerPort(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mockutils.NewMockHTTPClient(mockCtrl)
	sut := slb.New(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("slb/virtual-server/myvirtualserver/port/myport").Return(resp, nil)

	res, err := sut.VirtualServerPort.Get("myport", "myvirtualserver")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}
