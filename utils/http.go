package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

type httpClient struct {
	baseURL string
	debug   bool
	*http.Client
	*http.Header
}

type HTTPClient interface {
	GET(path string) (*Response, error)
	POST(path string, body interface{}) (*Response, error)
	PUT(path string, body interface{}) (*Response, error)
	DELETE(path string) (*Response, error)
	AddHeader(key string, value string)
	RemoveHeader(key string)
}

func NewHTTPClient(baseURL string, client *http.Client, header *http.Header, debug bool) HTTPClient {
	return &httpClient{baseURL: baseURL, Client: client, Header: header, debug: debug}
}

func (c *httpClient) GET(path string) (*Response, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s", c.baseURL, path),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	return &Response{res}, nil
}

func (c *httpClient) POST(path string, body interface{}) (*Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", c.baseURL, path),
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	return &Response{res}, nil
}

func (c *httpClient) PUT(path string, body interface{}) (*Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/%s", c.baseURL, path),
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	return &Response{res}, nil
}

func (c *httpClient) DELETE(path string) (*Response, error) {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/%s", c.baseURL, path),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	return &Response{res}, nil
}

func (c *httpClient) AddHeader(key string, value string) {
	c.Header.Add(key, value)
}

func (c *httpClient) RemoveHeader(key string) {
	c.Header.Del(key)
}

func (c *httpClient) do(req *http.Request) (*http.Response, error) {
	if c.debug {
		dump, _ := httputil.DumpRequest(req, true)
		fmt.Printf("---------------Request---------------\n%s\n", string(dump))
	}

	req.Header = *c.Header
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if c.debug {
		dump, _ := httputil.DumpResponse(res, true)
		fmt.Printf("---------------Response--------------\n%s\n", string(dump))
	}
	return res, nil
}

type Response struct {
	*http.Response
}

func (r *Response) HasError() bool {
	return r.StatusCode >= http.StatusBadRequest
}

func (r *Response) UnmarshalJSON(v interface{}) error {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return err
	}
	if err := json.Unmarshal(buf.Bytes(), v); err != nil {
		return err
	}
	return nil
}
