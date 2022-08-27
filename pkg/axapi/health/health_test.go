package health_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
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
