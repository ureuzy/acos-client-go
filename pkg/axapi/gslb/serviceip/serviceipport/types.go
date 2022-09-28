package serviceipport

import "github.com/ureuzy/acos-client-go/pkg/axapi/shared"

// Object Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_service_ip.html#service-ip-attributes

type Port struct {
	shared.AxaBase             `json:",inline"`
	PortNum                    int                     `json:"port-num,omitempty"`
	PortProto                  string                  `json:"port-proto,omitempty"`
	Action                     string                  `json:"action,omitempty"` //"enum":["enable", "disable"]
	HealthCheck                string                  `json:"health-check,omitempty"`
	HealthCheckFollowPort      int                     `json:"health-check-follow-port,omitempty"`
	FollowPortProtocol         string                  `json:"follow-port-protocol,omitempty"`
	HealthCheckProtocolDisable shared.Boolean          `json:"health-check-protocol-disable,omitempty"`
	HealthCheckDisable         shared.Boolean          `json:"health-check-disable,omitempty"`
	SamplingEnable             []shared.SamplingEnable `json:"sampling-enable,omitempty"` //"enum":["all","active","current"]
}
