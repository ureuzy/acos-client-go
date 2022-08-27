package serviceip

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/axapi/shared"
	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_service_ip.html#

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

// Object Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_service_ip.html#service-ip-attributes

type Object struct {
	shared.AxaBase             `json:",inline"`
	NodeName                   string                  `json:"node-name,omitempty"`
	IPv6Address                string                  `json:"ipv6-address,omitempty"`
	IPAddress                  string                  `json:"ip-address,omitempty"`
	Action                     string                  `json:"action,omitempty"` //"enum":["enable", "disable"]
	ExternalIP                 string                  `json:"external-ip,omitempty"`
	IPv6                       string                  `json:"ipv6,omitempty"`
	HealthCheck                string                  `json:"health-check,omitempty"`
	HealthCheckProtocolDisable bool                    `json:"health-check-protocol-disable,omitempty"`
	HealthCheckDisable         bool                    `json:"health-check-disable,omitempty"`
	SamplingEnable             []shared.SamplingEnable `json:"sampling-enable,omitempty"` //"enum":["all","hits","recent"]
	PortList                   []Port                  `json:"port-list,omitempty"`
}

type Port struct {
	shared.AxaBase             `json:",inline"`
	PortNum                    int                     `json:"port-num,omitempty"`
	PortProto                  string                  `json:"port-proto,omitempty"`
	Action                     string                  `json:"action,omitempty"` //"enum":["enable", "disable"]
	HealthCheck                string                  `json:"health-check,omitempty"`
	HealthCheckFollowPort      int                     `json:"health-check-follow-port,omitempty"`
	FollowPortProtocol         string                  `json:"follow-port-protocol,omitempty"`
	HealthCheckProtocolDisable bool                    `json:"health-check-protocol-disable,omitempty"`
	HealthCheckDisable         bool                    `json:"health-check-disable,omitempty"`
	SamplingEnable             []shared.SamplingEnable `json:"sampling-enable,omitempty"` //"enum":["all","active","current"]
}
