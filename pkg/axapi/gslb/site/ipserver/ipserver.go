package ipserver

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/site"
	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_zone_service.html

// /axapi/v3/gslb/site/{site-name}/ip-server/{ip-server-name}

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "site/%s/ip-server"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"ip-server-list"`
}

type Body struct {
	site.IPServer `json:"ip-server"`
}

type ListObjects []site.IPServer
