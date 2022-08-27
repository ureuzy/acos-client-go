package gslb_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	mockutils "github.com/ureuzy/acos-client-go/pkg/mocks/utils"
	"github.com/ureuzy/acos-client-go/utils"
)

func TestGetPolicy(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mockutils.NewMockHTTPClient(mockCtrl)
	sut := gslb.New(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("gslb/policy/mypolicy").Return(resp, nil)

	res, err := sut.Policy.Get("mypolicy")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetServiceIP(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mockutils.NewMockHTTPClient(mockCtrl)
	sut := gslb.New(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("gslb/service-ip/myip").Return(resp, nil)

	res, err := sut.ServiceIP.Get("myip")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetSite(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mockutils.NewMockHTTPClient(mockCtrl)
	sut := gslb.New(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("gslb/site/mysite").Return(resp, nil)

	res, err := sut.Site.Get("mysite")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestGetZone(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mockutils.NewMockHTTPClient(mockCtrl)
	sut := gslb.New(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("gslb/zone/myzone").Return(resp, nil)

	res, err := sut.Zone.Get("myzone")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}
