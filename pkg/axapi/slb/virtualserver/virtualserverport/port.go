package virtualserverport

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#port-specification

// /axapi/v3/slb/virtual-server/{name}/port/{port-number}+{protocol}

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "virtual-server/%s/port"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"port-list"`
}

type Body struct {
	Object `json:"port"`
}
