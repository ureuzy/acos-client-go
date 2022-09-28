package serviceip

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_service_ip.html#
// URI: /axapi/v3/gslb/service-ip/{node-name}

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "service-ip"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"service-ip-list"`
}

type Body struct {
	Object `json:"service-ip"`
}

type ListObjects []Object
