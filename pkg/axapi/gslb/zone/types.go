package zone

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/zone/zoneservice"
	"github.com/ureuzy/acos-client-go/pkg/axapi/shared"
)

// Object Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_zone.html#zone-attributes
type Object struct {
	shared.AxaBase  `json:",inline"`
	Name            string                    `json:"name,omitempty"`
	Disable         shared.Boolean            `json:"disable,omitempty"`
	Policy          string                    `json:"policy,omitempty"`
	Template        Template                  `json:"template,omitempty"`
	TTL             int                       `json:"ttl,omitempty"`
	UseServerTTL    shared.Boolean            `json:"use-server-ttl,omitempty"`
	DNSSOARecord    *zoneservice.DNSSOARecord `json:"dns-soa-record,omitempty"`
	SamplingEnable  []shared.SamplingEnable   `json:"sampling-enable,omitempty"` //"enum":[	"all",	"received-query",	"sent-response",	"proxy-mode-response",	"cache-mode-response",	"server-mode-response",	"sticky-mode-response",	"backup-mode-response"]
	DNSMXRecordList []zoneservice.DNSMXRecord `json:"dns-mx-record-list,omitempty"`
	DNSNSRecordList []zoneservice.DNSNSRecord `json:"dns-ns-record-list,omitempty"`
	ServiceList     []zoneservice.Service     `json:"service-list,omitempty"`
}

type Template struct {
	DNSSec string `json:"dnssec,omitempty"`
}
