package client

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/ureuzy/acos-client-go/pkg/axapi/auth"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb"
	"github.com/ureuzy/acos-client-go/utils"
)

type Client struct {
	conf Config
	http utils.HttpClient
	Auth auth.Operator
	Slb  *slb.Operator
}

type Config struct {
	Host  string
	User  string
	Pass  string
	Debug bool
}

type Option func(*http.Client)

func InsecureSkipVerify(b bool) Option {
	return func(client *http.Client) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: b},
		}
	}
}

func New(conf Config, options ...Option) (*Client, error) {
	httpClient := &http.Client{}
	for _, opt := range options {
		opt(httpClient)
	}

	baseUrl := fmt.Sprintf("https://%s/axapi/v3", conf.Host)
	client := generateClient(utils.NewHttpClient(baseUrl, httpClient, &http.Header{}, conf.Debug), conf)
	client.http.AddHeader("Content-Type", "application/json")

	if err := client.Authenticate(); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Authenticate() error {
	if res, err := c.Auth.Request(&auth.Request{Credentials: auth.Credentials{
		Username: c.conf.User,
		Password: c.conf.Pass,
	}}); err != nil {
		return err
	} else {
		c.http.AddHeader("Authorization", fmt.Sprintf("A10 %s", res.Signature))
	}
	return nil
}

func generateClient(client utils.HttpClient, conf Config) *Client {
	return &Client{conf, client, auth.New(client), slb.New(client)}
}
