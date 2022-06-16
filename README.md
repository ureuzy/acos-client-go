# acos-client-go

[![Go Reference](https://pkg.go.dev/badge/github.com/masanetes/acos-client-go.svg)](https://pkg.go.dev/github.com/masanetes/acos-client-go)

A simple go client for [a10 networks](https://www.a10networks.com/)' aXAPI

[ACOS 4.1.4-P1 aXAPIv3 Reference Document](https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/index.html#)

## Documentation

[godoc](https://pkg.go.dev/github.com/masanetes/acos-client-go)

## Usage

import client 

```go
import "github.com/masanetes/acos-client-go"
```

New client

```go
config := client.Config{Host: "<HOST>", User: "<USER>", Pass: "<PASS>", Debug: false}
c, _ := client.New(&config, client.InsecureSkipVerify(true))

// Get virtual server
vs, _ := c.Slb.GetVirtualServer("masanetes-sample-virtualserver")
fmt.Println(vs.Name, vs.IPAddress)
```


