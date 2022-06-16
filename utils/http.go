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
	baseUrl string
	debug   bool
	*http.Client
	*http.Header
}

type HttpClient interface {
	GET(path string) (*response, error)
	POST(path string, body []byte) (*response, error)
	PUT(path string, body []byte) (*response, error)
	DELETE(path string) (*response, error)
	AddHeader(key string, value string)
}

func NewHttpClient(baseUrl string, client *http.Client, header *http.Header, debug bool) HttpClient {
	return &httpClient{baseUrl: baseUrl, Client: client, Header: header, debug: debug}
}

func (c *httpClient) GET(path string) (*response, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s", c.baseUrl, path),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	return &response{res}, nil
}

func (c *httpClient) POST(path string, body []byte) (*response, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", c.baseUrl, path),
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	return &response{res}, nil
}

func (c *httpClient) PUT(path string, body []byte) (*response, error) {
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/%s", c.baseUrl, path),
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	return &response{res}, nil
}

func (c *httpClient) DELETE(path string) (*response, error) {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/%s", c.baseUrl, path),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	return &response{res}, nil
}

func (c *httpClient) AddHeader(key string, value string) {
	c.Header.Add(key, value)
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

type response struct {
	*http.Response
}

func (r *response) UnmarshalJson(v interface{}) error {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return err
	}
	if err := json.Unmarshal(buf.Bytes(), v); err != nil {
		return err
	}
	return nil
}
