package testutils

import (
	"io"
	"net/http"
	"strings"

	"github.com/ureuzy/acos-client-go/utils"
)

func ResponseOK() *utils.Response {
	return &utils.Response{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("{}")),
		},
	}
}

func ResponseNotFound() *utils.Response {
	return &utils.Response{
		Response: &http.Response{
			Status:     "404 not found",
			StatusCode: http.StatusNotFound,
			Body:       io.NopCloser(strings.NewReader("{}")),
		},
	}
}
