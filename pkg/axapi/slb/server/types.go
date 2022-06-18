package server

type ListObject struct {
	ServerList `json:"server-list"`
}

type Object struct {
	Server `json:"server"`
}

// Server Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#server-attributes
type Server struct {
	Action             string              `json:"action,omitempty"`
	AlternateServer    *AlternateServer    `json:"alternate-server,omitempty"`
	ConnLimit          int                 `json:"conn-limit,omitempty"`
	ConnResume         int                 `json:"conn-resume,omitempty"`
	ExtendedStats      int                 `json:"extended-stats,omitempty"`
	ExternalIP         string              `json:"external-ip,omitempty"`
	FqdnName           string              `json:"fqdn-name,omitempty"`
	HealthCheck        string              `json:"health-check,omitempty"`
	HealthCheckDisable int                 `json:"health-check-disable,omitempty"`
	Host               string              `json:"host,omitempty"`
	IPv6               string              `json:"ipv6,omitempty"`
	Name               string              `json:"name,omitempty"`
	NoLogging          int                 `json:"no-logging,omitempty"`
	PortList           *PortList           `json:"port-list,omitempty"`
	SamplingEnable     *SamplingEnableList `json:"sampling-enable,omitempty"`
	ServerIPv6Addr     string              `json:"server-ipv6-addr,omitempty"`
	SlowStart          int                 `json:"slow-start,omitempty"`
	SpoofingCache      int                 `json:"spoofing-cache,omitempty"`
	StatsDataAction    string              `json:"stats-data-action,omitempty"`
	TemplateServer     string              `json:"template-server,omitempty"`
	UserTag            string              `json:"user-tag,omitempty"`
	UUID               string              `json:"uuid,omitempty"`
	Weight             int                 `json:"weight,omitempty"`
}
type ServerList []Server

// AlternateServer Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#alternate-server
type AlternateServer struct {
	Alternate     int    `json:"alternate,omitempty"`
	AlternateName string `json:"alternate-name,omitempty"`
}
type AlternateServerList []AlternateServer

// Port Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#port-list
type Port struct {
	Action                string              `json:"action,omitempty"`
	AlternatePort         []string            `json:"alternate-port,omitempty"`
	AuthCfg               *AuthCfg            `json:"auth-cfg,omitempty"`
	ConnLimit             int                 `json:"conn-limit,omitempty"`
	ConnResume            int                 `json:"conn-resume,omitempty"`
	ExtendedStats         int                 `json:"extended-stats,omitempty"`
	FollowPortProtocol    string              `json:"follow-port-protocol,omitempty"`
	HealthCheck           string              `json:"health-check,omitempty"`
	HealthCheckDisable    int                 `json:"health-check-disable,omitempty"`
	HealthCheckFollowPort int                 `json:"health-check-follow-port,omitempty"`
	NoLogging             int                 `json:"no-logging,omitempty"`
	NoSSL                 int                 `json:"no-ssl,omitempty"`
	PortNumber            int                 `json:"port-number,omitempty"`
	Protocol              string              `json:"protocol,omitempty"`
	Range                 int                 `json:"range,omitempty"`
	SamplingEnable        *SamplingEnableList `json:"sampling-enable,omitempty"`
	ServerIPv6Addr        string              `json:"server-ipv6-addr,omitempty"`
	SlowStart             int                 `json:"slow-start,omitempty"`
	SpoofingCache         int                 `json:"spoofing-cache,omitempty"`
	StatsDataAction       string              `json:"stats-data-action,omitempty"`
	TemplateServer        string              `json:"template-server,omitempty"`
	UserTag               string              `json:"user-tag,omitempty"`
	UUID                  string              `json:"uuid,omitempty"`
	Weight                int                 `json:"weight,omitempty"`
}
type PortList []Port

// AuthCfg Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#port-list-auth-cfg
type AuthCfg struct {
	ServicePrincipalName string `json:"service-principal-name,omitempty"`
}

// SamplingEnable Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_server.html#sampling-enable
type SamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SamplingEnableList []SamplingEnable
