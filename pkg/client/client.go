package client

import (
	"fmt"
	"net/http"

	ap "github.com/ureuzy/acos-client-go/pkg/axapi/activepartition"
	"github.com/ureuzy/acos-client-go/pkg/axapi/auth"
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb"
	"github.com/ureuzy/acos-client-go/pkg/axapi/health"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb"
	"github.com/ureuzy/acos-client-go/utils"
)

type Client struct {
	conf            Config
	http            utils.HTTPClient
	Auth            auth.Operator
	ActivePartition ap.Operator
	Health          *health.Operator
	Slb             *slb.Operator
	Gslb            *gslb.Operator
}

type Config struct {
	Host  string
	User  string
	Pass  string
	Debug bool
}

type Option func(*http.Client)

func Instance(conf Config, options ...Option) *Client {
	httpClient := &http.Client{}
	for _, opt := range options {
		opt(httpClient)
	}

	baseURL := fmt.Sprintf("https://%s/axapi/v3", conf.Host)
	client := generateClient(utils.NewHTTPClient(baseURL, httpClient, &http.Header{}, conf.Debug), conf)
	client.http.AddHeader("Content-Type", "application/json")

	return client
}

func New(conf Config, options ...Option) (*Client, error) {
	httpClient := &http.Client{}
	for _, opt := range options {
		opt(httpClient)
	}

	baseURL := fmt.Sprintf("https://%s/axapi/v3", conf.Host)
	client := generateClient(utils.NewHTTPClient(baseURL, httpClient, &http.Header{}, conf.Debug), conf)
	client.http.AddHeader("Content-Type", "application/json")

	if err := client.Authenticate(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Authenticate() error {
	res, err := c.Auth.Login(&auth.Request{Credentials: auth.Credentials{
		Username: c.conf.User,
		Password: c.conf.Pass,
	}})
	if err != nil {
		return err
	}
	c.http.AddHeader("Authorization", fmt.Sprintf("A10 %s", res.Signature))
	return nil
}

func (c *Client) Logoff() error {
	err := c.Auth.Logoff()
	if err != nil {
		return err
	}
	c.http.RemoveHeader("Authorization")
	return nil
}

func generateClient(client utils.HTTPClient, conf Config) *Client {
	return &Client{
		conf:            conf,
		http:            client,
		Auth:            auth.New(client),
		ActivePartition: ap.New(client),
		Health:          health.New(client),
		Slb:             slb.New(client),
		Gslb:            gslb.New(client),
	}
}
