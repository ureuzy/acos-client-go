package zoneservice

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_zone_service.html
// URI: /axapi/v3/gslb/zone/{name}/service/{service-port}+{service-name}

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "zone/%s/service"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"service-list"`
}

type Body struct {
	Service `json:"service"`
}

type ListObjects []Service
