package serviceip

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/serviceip/serviceipport"
	"github.com/ureuzy/acos-client-go/pkg/axapi/shared"
)

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
	HealthCheckProtocolDisable shared.Boolean          `json:"health-check-protocol-disable,omitempty"`
	HealthCheckDisable         shared.Boolean          `json:"health-check-disable,omitempty"`
	SamplingEnable             []shared.SamplingEnable `json:"sampling-enable,omitempty"` //"enum":["all","hits","recent"]
	PortList                   []serviceipport.Port    `json:"port-list,omitempty"`
}
