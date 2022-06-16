package client

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/masanetes/acos-client-go/pkg/axapi/auth"
	"github.com/masanetes/acos-client-go/pkg/axapi/slb"
	"github.com/masanetes/acos-client-go/utils"
)

type Client struct {
	Auth auth.Operator
	Slb  slb.Operator
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

func New(conf *Config, options ...Option) (*Client, error) {
	httpClient := &http.Client{}
	for _, opt := range options {
		opt(httpClient)
	}

	baseUrl := fmt.Sprintf("https://%s/axapi/v3", conf.Host)
	c := utils.NewHttpClient(baseUrl, httpClient, &http.Header{}, conf.Debug)
	c.AddHeader("Content-Type", "application/json")

	client := generateClient(c)
	res, err := client.Auth.Request(&auth.Request{Credentials: auth.Credentials{
		Username: conf.User,
		Password: conf.Pass,
	}})
	if err != nil {
		return &Client{}, err
	}
	c.AddHeader("Authorization", fmt.Sprintf("A10 %s", res.Signature))

	return client, err
}

func generateClient(client utils.HttpClient) *Client {
	return &Client{
		auth.New(client),
		slb.New(client),
	}
}
