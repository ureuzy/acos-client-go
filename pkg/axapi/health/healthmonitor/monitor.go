package healthmonitor

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/health_monitor.html

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "monitor"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"monitor-list,omitempty"`
}

type Body struct {
	Object `json:"monitor,omitempty"`
}

type ListObjects []Object
