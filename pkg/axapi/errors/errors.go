package errors

import (
	"errors"

	"github.com/masanetes/acos-client-go/utils"
)

type ResponseBody struct {
	Response `json:"response"`
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

func EmptyStringError(s string) error {
	if s == "" {
		return errors.New("must be specified")
	}
	return nil
}

func Handle(response *utils.Response) error {
	var errResponse ResponseBody
	if err := response.UnmarshalJson(&errResponse); err != nil {
		return err
	}
	return &errResponse
}
