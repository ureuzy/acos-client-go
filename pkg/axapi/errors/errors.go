package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ureuzy/acos-client-go/utils"
)

type ResponseBody struct {
	Status     string
	StatusCode int
	Response   `json:"response"`
}

type Response struct {
	Status string `json:"status"`
	Err    `json:"err"`
}

type Err struct {
	Code     int    `json:"code"`
	From     string `json:"from"`
	Msg      string `json:"msg"`
	Location string `json:"location"`
}

func (r *ResponseBody) Error() string {
	return r.Msg
}

func (r *ResponseBody) Unwrap() error {
	return r
}

func ArgsMismatchError(pathf string, vals []string) error {
	cnt := strings.Count(pathf, "%")
	if cnt != len(vals) {
		return fmt.Errorf("wrong number of arguments: required: %d actual %d", cnt, len(vals))
	}
	return nil
}

func EmptyStringError(s string) error {
	if s == "" {
		return errors.New("identifier must be specified")
	}
	return nil
}

// Handle is a function that returns an error response from an ACOS device as an error interface.
func Handle(response *utils.Response) error {
	var errResponse ResponseBody
	if err := response.UnmarshalJSON(&errResponse); err != nil {
		return err
	}
	errResponse.Status = response.Status
	errResponse.StatusCode = response.StatusCode
	return &errResponse
}
