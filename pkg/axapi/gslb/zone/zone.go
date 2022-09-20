package zone

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/axapi/shared"
	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_zone.html

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "zone"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"zone-list"`
}

type Body struct {
	Object `json:"zone"`
}

type ListObjects []Object

// Object Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_zone.html#zone-attributes
type Object struct {
	shared.AxaBase  `json:",inline"`
	Name            string                  `json:"name,omitempty"`
	Disable         shared.Boolean          `json:"disable,omitempty"`
	Policy          string                  `json:"policy,omitempty"`
	Template        Template                `json:"template,omitempty"`
	TTL             int                     `json:"ttl,omitempty"`
	UseServerTTL    shared.Boolean          `json:"use-server-ttl,omitempty"`
	DNSSOARecord    DNSSOARecord            `json:"dns-soa-record,omitempty"`
	SamplingEnable  []shared.SamplingEnable `json:"sampling-enable,omitempty"` //"enum":[	"all",	"received-query",	"sent-response",	"proxy-mode-response",	"cache-mode-response",	"server-mode-response",	"sticky-mode-response",	"backup-mode-response"]
	DNSMXRecordList []DNSMXRecord           `json:"dns-mx-record-list,omitempty"`
	DNSNSRecordList []DNSNSRecord           `json:"dns-ns-record-list,omitempty"`
	ServiceList     []Service               `json:"service-list,omitempty"`
}

type Template struct {
	DNSSec string `json:"dnssec,omitempty"`
}

type DNSSOARecord struct {
	SOAName   string `json:"soa-name,omitempty"`
	Mail      string `json:"mail,omitempty"`
	Expire    int    `json:"expire,omitempty"`
	Refresh   int    `json:"refresh,omitempty"`
	Retry     int    `json:"retry,omitempty"`
	Serial    int    `json:"serial,omitempty"`
	SOATTL    int    `json:"soa-ttl,omitempty"`
	External  string `json:"external,omitempty"`
	ExMail    string `json:"ex-mail,omitempty"`
	ExEpire   int    `json:"ex-expire,omitempty"`
	ExRefresh int    `json:"ex-refresh,omitempty"`
	ExRetry   int    `json:"ex-retry,omitempty"`
	ExSerial  int    `json:"ex-serial,omitempty"`
	ExSoaTTL  int    `json:"ex-soa-ttl,omitempty"`
}

type Service struct {
	shared.AxaBase     `json:",inline"`
	ServicePort        int                     `json:"service-port,omitempty"`
	ServiceName        string                  `json:"service-name,omitempty"`
	Action             string                  `json:"action,omitempty"`       // "enum":[	"drop",	"forward",	"ignore",	"reject"]
	ForwardType        string                  `json:"forward-type,omitempty"` // "enum":[	"both",	"query",	"response"]
	Disable            shared.Boolean          `json:"disable,omitempty"`
	HealthCheckGateway string                  `json:"health-check-gateway,omitempty"` //"enum":[	"enable",	"disable"]
	HealthCheckPort    []HealthCheckPort       `json:"health-check-port,omitempty"`
	Policy             string                  `json:"policy,omitempty"`
	SamplingEnable     []shared.SamplingEnable `json:"sampling-enable,omitempty"` // "enum":[	"all",	"received-query",	"sent-response",	"proxy-mode-response",	"cache-mode-response",	"server-mode-response",	"sticky-mode-response",	"backup-mode-response"]
	DNSARecord         DNSARecord              `json:"dns-a-record,omitempty"`
	DNSCNAMERecordList []DNSCnameRecord        `json:"dns-cname-record-list,omitempty"`
	DNSMXRecordList    []DNSMXRecord           `json:"dns-mx-record-list,omitempty"`
	DNSNSRecordList    []DNSMXRecord           `json:"dns-ns-record-list,omitempty"`
	DNSPtrRecordList   []TodoRecord            `json:"dns-ptr-record-list,omitempty"`
	DNSSrvRecordList   []TodoRecord            `json:"dns-srv-record-list,omitempty"`
	DNSNaptrRecordList []TodoRecord            `json:"dns-naptr-record-list,omitempty"`
	DNSTxtRecordList   []TodoRecord            `json:"dns-txt-record-list,omitempty"`
	DNSRecordList      []TodoRecord            `json:"dns-record-list,omitempty"`
	GeoLocationList    []TodoRecord            `json:"geo-location-list,omitempty"`
}

type DNSMXRecord struct {
	shared.AxaBase `json:",inline"`
}

type DNSNSRecord struct {
	shared.AxaBase `json:",inline"`
}

type TodoRecord struct {
	// TODO
}

type HealthCheckPort struct {
	HealthCheckPort int `json:"health-check-port,omitempty"`
}

type DNSARecord struct {
	DNSARecordSrvList  []DNSARecordSrv  `json:"dns-a-record-srv-list,omitempty"`
	DNSARecordIPv4List []DNSARecordIPv4 `json:"dns-a-record-ipv4-list,omitempty"`
	DNSARecordIPv6List []DNSARecordIPv6 `json:"dns-a-record-ipv6-list,omitempty"`
}

type DNSCnameRecord struct {
	AliasName       string                  `json:"alias-name,omitempty"`
	AdminPreference shared.Boolean          `json:"admin-preference,omitempty"`
	Weight          int                     `json:"weight,omitempty"`
	AsBackup        shared.Boolean          `json:"as-backup,omitempty"`
	SamplingEnable  []shared.SamplingEnable `json:"sampling-enable,omitempty"` // "enum":[	"all",	"cname-hits"]
}

type DNSARecordSrv struct {
	shared.AxaBase
	SvrName   string         `json:"svrname,omitempty"`
	NoResp    shared.Boolean `json:"no-resp,omitempty"`
	AsBackup  shared.Boolean `json:"as-backup,omitempty"`
	Weight    int            `json:"weight,omitempty"`
	TTL       int            `json:"ttl,omitempty"`
	AsReplace shared.Boolean `json:"as-replace,omitempty"`
	Disable   shared.Boolean `json:"disable,omitempty"`
	Static    shared.Boolean `json:"static,omitempty"`
	AdminIP   int            `json:"admin-ip,omitempty"`
}

type DNSARecordIPv4 struct {
	shared.AxaBase
	DNSARecordIP string         `json:"dns-a-record-ip,omitempty"`
	NoResp       shared.Boolean `json:"no-resp,omitempty"`
	AsBackup     shared.Boolean `json:"as-backup,omitempty"`
	Weight       int            `json:"weight,omitempty"`
	TTL          int            `json:"ttl,omitempty"`
	AsReplace    shared.Boolean `json:"as-replace,omitempty"`
	Disable      shared.Boolean `json:"disable,omitempty"`
	Static       shared.Boolean `json:"static,omitempty"`
	AdminIP      int            `json:"admin-ip,omitempty"`
}

type DNSARecordIPv6 struct {
	shared.AxaBase
	DNSARecordIP string         `json:"dns-a-record-ipv6,omitempty"`
	NoResp       shared.Boolean `json:"no-resp,omitempty"`
	AsBackup     shared.Boolean `json:"as-backup,omitempty"`
	Weight       int            `json:"weight,omitempty"`
	TTL          int            `json:"ttl,omitempty"`
	AsReplace    shared.Boolean `json:"as-replace,omitempty"`
	Disable      shared.Boolean `json:"disable,omitempty"`
	Static       shared.Boolean `json:"static,omitempty"`
	AdminIP      int            `json:"admin-ip,omitempty"`
}
