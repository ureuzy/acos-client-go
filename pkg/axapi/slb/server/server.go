package server

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#server-specification

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "server"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"server-list"`
}

type Body struct {
	Object `json:"server"`
}
