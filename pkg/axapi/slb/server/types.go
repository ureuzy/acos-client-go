package server

import "github.com/ureuzy/acos-client-go/pkg/axapi/shared"

// Object Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#server-attributes
type Object struct {
	shared.AxaBase     `json:",inline"`
	Action             string                 `json:"action,omitempty"`
	AlternateServer    *AlternateServer       `json:"alternate-server,omitempty"`
	ConnLimit          int                    `json:"conn-limit,omitempty"`
	ConnResume         int                    `json:"conn-resume,omitempty"`
	ExtendedStats      shared.Boolean         `json:"extended-stats,omitempty"`
	ExternalIP         string                 `json:"external-ip,omitempty"`
	FqdnName           string                 `json:"fqdn-name,omitempty"`
	HealthCheck        string                 `json:"health-check,omitempty"`
	HealthCheckDisable shared.Boolean         `json:"health-check-disable,omitempty"`
	Host               string                 `json:"host,omitempty"`
	IPv6               string                 `json:"ipv6,omitempty"`
	Name               string                 `json:"name,omitempty"`
	NoLogging          shared.Boolean         `json:"no-logging,omitempty"`
	PortList           *PortList              `json:"port-list,omitempty"`
	SamplingEnable     *shared.SamplingEnable `json:"sampling-enable,omitempty"`
	ServerIPv6Addr     string                 `json:"server-ipv6-addr,omitempty"`
	SlowStart          shared.Boolean         `json:"slow-start,omitempty"`
	SpoofingCache      shared.Boolean         `json:"spoofing-cache,omitempty"`
	StatsDataAction    string                 `json:"stats-data-action,omitempty"`
	TemplateServer     string                 `json:"template-server,omitempty"`
	Weight             int                    `json:"weight,omitempty"`
}
type ListObjects []Object

// AlternateServer Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#alternate-server
type AlternateServer struct {
	Alternate     int    `json:"alternate,omitempty"`
	AlternateName string `json:"alternate-name,omitempty"`
}
type AlternateServerList []AlternateServer

// Port Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#port-list
type Port struct {
	shared.AxaBase        `json:",inline"`
	Action                string                 `json:"action,omitempty"`
	AlternatePort         []string               `json:"alternate-port,omitempty"`
	AuthCfg               *AuthCfg               `json:"auth-cfg,omitempty"`
	ConnLimit             int                    `json:"conn-limit,omitempty"`
	ConnResume            int                    `json:"conn-resume,omitempty"`
	ExtendedStats         shared.Boolean         `json:"extended-stats,omitempty"`
	FollowPortProtocol    string                 `json:"follow-port-protocol,omitempty"`
	HealthCheck           string                 `json:"health-check,omitempty"`
	HealthCheckDisable    shared.Boolean         `json:"health-check-disable,omitempty"`
	HealthCheckFollowPort int                    `json:"health-check-follow-port,omitempty"`
	NoLogging             shared.Boolean         `json:"no-logging,omitempty"`
	NoSSL                 shared.Boolean         `json:"no-ssl,omitempty"`
	PortNumber            int                    `json:"port-number,omitempty"`
	Protocol              string                 `json:"protocol,omitempty"`
	Range                 int                    `json:"range,omitempty"`
	SamplingEnable        *shared.SamplingEnable `json:"sampling-enable,omitempty"`
	ServerIPv6Addr        string                 `json:"server-ipv6-addr,omitempty"`
	SlowStart             int                    `json:"slow-start,omitempty"`
	SpoofingCache         int                    `json:"spoofing-cache,omitempty"`
	StatsDataAction       string                 `json:"stats-data-action,omitempty"`
	TemplateServer        string                 `json:"template-server,omitempty"`
	Weight                int                    `json:"weight,omitempty"`
}
type PortList []Port

// AuthCfg Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#port-list-auth-cfg
type AuthCfg struct {
	ServicePrincipalName string `json:"service-principal-name,omitempty"`
}
