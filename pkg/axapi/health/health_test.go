package health_test

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	a10errors "github.com/ureuzy/acos-client-go/pkg/axapi/errors"
	"github.com/ureuzy/acos-client-go/pkg/axapi/health"
	mockutils "github.com/ureuzy/acos-client-go/pkg/mocks/utils"
	"github.com/ureuzy/acos-client-go/utils"
)

func TestGetMonitor(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mockutils.NewMockHTTPClient(mockCtrl)
	sut := health.New(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("health/monitor/resource1").Return(resp, nil)

	res, err := sut.Montitor.Get("resource1")
	Ω(err).ShouldNot(HaveOccurred())
	Ω(res).ShouldNot(BeNil())
}

func TestNotFound(t *testing.T) {
	RegisterTestingT(t)

	mockCtrl := gomock.NewController(t)
	httpc := mockutils.NewMockHTTPClient(mockCtrl)
	sut := health.New(httpc)

	body := io.NopCloser(strings.NewReader("{}"))

	resp := &utils.Response{
		Response: &http.Response{
			Status:     "404 not found",
			StatusCode: http.StatusNotFound,
			Body:       body,
		},
	}

	httpc.EXPECT().GET("health/monitor/resource1").Return(resp, nil)

	_, err := sut.Montitor.Get("resource1")

	Ω(err).Should(HaveOccurred())

	var errUnwrapped *a10errors.ResponseBody
	ok := errors.As(err, &errUnwrapped)
	Ω(ok).Should(BeTrue())

	Ω(errUnwrapped.StatusCode).Should(BeIdenticalTo(http.StatusNotFound))
}
