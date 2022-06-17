# acos-client-go

[![Go Reference](https://pkg.go.dev/badge/github.com/masanetes/acos-client-go.svg)](https://pkg.go.dev/github.com/masanetes/acos-client-go)

A simple go client for [a10 networks](https://www.a10networks.com/)' aXAPI

[ACOS 4.1.4-P1 aXAPIv3 Reference Document](https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/index.html#)

## Documentation

[godoc](https://pkg.go.dev/github.com/masanetes/acos-client-go)

## Install

```
go get -u github.com/masanetes/acos-client-go
```

## Usage

Import

```go
import "github.com/masanetes/acos-client-go"
```

New client

```go
config := client.Config{Host: "<HOST>", User: "<USER>", Pass: "<PASS>", Debug: false}
c, _ := client.New(config)

// Get virtual server
vs, _ := c.Slb.GetVirtualServer("masanetes-sample-virtualserver")
fmt.Println(vs.Name, vs.IPAddress)
```

## Optional Parameter

When creating a client, you can accept optional arguments and customize the http client

```go
customHTTPClient := func(client *http.Client) {
	client = &http.Client{} 
}
client.New(config, customHTTPClient)
```

There are also several options available that are likely to be used more frequently and can be used.

**TLS insecure skip verify**

```go
client.New(config, client.InsecureSkipVerify(true))
```

## Authenticate

A token is issued when the client is created, and is added to the header as a bearer token, but if the token expires, the client can re-authenticate.

```go
client.Authenticate()
```
