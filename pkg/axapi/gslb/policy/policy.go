package policy

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_policy.html

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "policy"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"policy-list"`
}

type Body struct {
	Object `json:"policy"`
}

type ListObjects []Object
