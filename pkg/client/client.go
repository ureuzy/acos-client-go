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

// Client is an ACOS client structure. It has each resource interface and HTTP client.
type Client struct {
	conf            Config
	http            utils.HTTPClient
	Auth            auth.Operator
	ActivePartition ap.Operator
	Health          *health.Operator
	Slb             *slb.Operator
	Gslb            *gslb.Operator
}

// Config is information required by the ACOS client.
// It is a structure containing host and credential information.
type Config struct {
	Host  string
	User  string
	Pass  string
	Debug bool
}

// Option is a function to customize the HTTP client included in the ACOS client
type Option func(*http.Client)

// NewInstance creates a new unauthenticated ACOS client.
// To authenticate, call NewAuthenticated instead of NewInstance
// or use the Authenticate method on the client instance
func NewInstance(conf Config, options ...Option) *Client {
	httpClient := &http.Client{}
	for _, opt := range options {
		opt(httpClient)
	}

	baseURL := fmt.Sprintf("https://%s/axapi/v3", conf.Host)
	client := generateClient(utils.NewHTTPClient(baseURL, httpClient, &http.Header{}, conf.Debug), conf)
	client.http.AddHeader("Content-Type", "application/json")

	return client
}

// NewAuthenticated creates a new authenticated ACOS client.
// If the token has expired, it can be re-authenticated by calling the Authenticate method
func NewAuthenticated(conf Config, options ...Option) (*Client, error) {
	client := NewInstance(conf, options...)

	if err := client.Authenticate(); err != nil {
		return nil, err
	}

	return client, nil
}

// Authenticate performs the login process to the ACOS device
// and sets the obtained Authorization token to the HTTP client
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

// Logoff removes the Authorization header from the HTTP client
// after the logoff process from the ACOS device
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
