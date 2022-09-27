package port

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/axapi/shared"
	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#port-specification

// /axapi/v3/slb/virtual-server/{name}/port/{port-number}+{protocol}

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "virtual-server/%s/port"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"port-list"`
}

type Body struct {
	Object `json:"port"`
}

// Object Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#port-attributes
type Object struct {
	shared.AxaBase                   `json:",inline"`
	ACLIDList                        *ACLIDList             `json:"acl-id-list,omitempty"`
	ACLNameList                      *ACLNameList           `json:"acl-name-list,omitempty"`
	Action                           string                 `json:"action,omitempty"`
	AflexScripts                     *AflexScripts          `json:"aflex-scripts,omitempty"`
	AltProtocol1                     string                 `json:"alt-protocol1,omitempty"`
	AltProtocol2                     string                 `json:"alt-protocol2,omitempty"`
	AlternatePort                    shared.Boolean         `json:"alternate-port,omitempty"`
	AlternatePortNumber              int                    `json:"alternate-port-number,omitempty"`
	AuthCfg                          *AuthCfg               `json:"auth-cfg,omitempty"`
	Auto                             shared.Boolean         `json:"auto,omitempty"`
	ClientIPStickyNat                shared.Boolean         `json:"clientip-sticky-nat,omitempty"`
	ConnLimit                        int                    `json:"conn-limit,omitempty"`
	DefSelectionIfPrefFailed         string                 `json:"def-selection-if-pref-failed,omitempty"`
	EnablePlayerIDCheck              shared.Boolean         `json:"enable-playerid-check,omitempty"`
	EthFwd                           int                    `json:"eth-fwd,omitempty"`
	EthRev                           int                    `json:"eth-rev,omitempty"`
	Expand                           shared.Boolean         `json:"expand,omitempty"`
	ExtendedStats                    shared.Boolean         `json:"extended-stats,omitempty"`
	ForceRoutingMode                 shared.Boolean         `json:"force-routing-mode,omitempty"`
	GslbEnable                       shared.Boolean         `json:"gslb-enable,omitempty"`
	HAConnMirror                     shared.Boolean         `json:"ha-conn-mirror,omitempty"`
	IPInIP                           shared.Boolean         `json:"ipinip,omitempty"`
	L7HardwareAssist                 shared.Boolean         `json:"l7-hardware-assist,omitempty"`
	MessageSwitching                 shared.Boolean         `json:"message-switching,omitempty"`
	Name                             string                 `json:"name,omitempty"`
	NoAutoUpOnAflex                  shared.Boolean         `json:"no-auto-up-on-aflex,omitempty"`
	NoDestNat                        shared.Boolean         `json:"no-dest-nat,omitempty"`
	NoLogging                        shared.Boolean         `json:"no-logging,omitempty"`
	OnSyn                            shared.Boolean         `json:"on-syn,omitempty"`
	PersistType                      string                 `json:"persist-type,omitempty"`
	Pool                             string                 `json:"pool,omitempty"`
	PortNumber                       int                    `json:"port-number,omitempty"`
	PortTranslation                  shared.Boolean         `json:"port-translation,omitempty"`
	Precedence                       shared.Boolean         `json:"precedence,omitempty"`
	Protocol                         string                 `json:"protocol,omitempty"`
	Range                            int                    `json:"range,omitempty"`
	Rate                             int                    `json:"rate,omitempty"`
	RedirectToHTTPS                  shared.Boolean         `json:"redirect-to-https,omitempty"`
	ReqFail                          shared.Boolean         `json:"req-fail,omitempty"`
	Reset                            shared.Boolean         `json:"reset,omitempty"`
	ResetOnServerSelectionFail       shared.Boolean         `json:"reset-on-server-selection-fail,omitempty"`
	RtpSipCallIDMatch                shared.Boolean         `json:"rtp-sip-call-id-match,omitempty"`
	SamplingEnable                   *shared.SamplingEnable `json:"sampling-enable,omitempty"`
	ScaleOutBucketCount              int                    `json:"scaleout-bucket-count,omitempty"`
	ScaleOutDeviceGroup              int                    `json:"scaleout-device-group,omitempty"`
	Secs                             int                    `json:"secs,omitempty"`
	ServSelFail                      shared.Boolean         `json:"serv-sel-fail,omitempty"`
	ServiceGroup                     string                 `json:"service-group,omitempty"`
	SharedPartitionClientSSLTemplate shared.Boolean         `json:"shared-partition-client-ssl-template,omitempty"`
	SharedPartitionHTTPTemplate      shared.Boolean         `json:"shared-partition-http-template,omitempty"`
	SharedPartitionServerSSLTemplate shared.Boolean         `json:"shared-partition-server-ssl-template,omitempty"`
	SharedPartitionTCP               shared.Boolean         `json:"shared-partition-tcp,omitempty"`
	SharedPartitionUDP               shared.Boolean         `json:"shared-partition-udp,omitempty"`
	SkipRevHash                      shared.Boolean         `json:"skip-rev-hash,omitempty"`
	SNatOnVip                        shared.Boolean         `json:"snat-on-vip,omitempty"`
	StatsDataAction                  string                 `json:"stats-data-action,omitempty"`
	SupportHTTP2                     shared.Boolean         `json:"support-http2,omitempty"`
	SynCookie                        shared.Boolean         `json:"syn-cookie,omitempty"`
	TemplateCache                    string                 `json:"template-cache,omitempty"`
	TemplateClientSSL                string                 `json:"template-client-ssl,omitempty"`
	TemplateClientSSLShared          string                 `json:"template-client-ssl-shared,omitempty"`
	TemplateConnectionReuse          string                 `json:"template-connection-reuse,omitempty"`
	TemplateDBLB                     string                 `json:"template-dblb,omitempty"`
	TemplateDiameter                 string                 `json:"template-diameter,omitempty"`
	TemplateDNS                      string                 `json:"template-dns,omitempty"`
	TemplateDynamicService           string                 `json:"template-dynamic-service,omitempty"`
	TemplateExternalService          string                 `json:"template-external-service,omitempty"`
	TemplateFileInspection           string                 `json:"template-file-inspection,omitempty"`
	TemplateFix                      string                 `json:"template-fix,omitempty"`
	TemplateFTP                      string                 `json:"template-ftp,omitempty"`
	TemplateHTTP                     string                 `json:"template-http,omitempty"`
	TemplateHTTPPolicy               string                 `json:"template-http-policy,omitempty"`
	TemplateHTTPShared               string                 `json:"template-http-shared,omitempty"`
	TemplateImapPop3                 string                 `json:"template-imap-pop3,omitempty"`
	TemplatePersistCookie            string                 `json:"template-persist-cookie,omitempty"`
	TemplatePersistDestinationIP     string                 `json:"template-persist-destination-ip,omitempty"`
	TemplatePersistSourceIP          string                 `json:"template-persist-source-ip,omitempty"`
	TemplatePersistSSLSid            string                 `json:"template-persist-ssl-sid,omitempty"`
	TemplatePolicy                   string                 `json:"template-policy,omitempty"`
	TemplateReqmodIcap               string                 `json:"template-reqmod-icap,omitempty"`
	TemplateRespmodIcap              string                 `json:"template-respmod-icap,omitempty"`
	TemplateScaleout                 string                 `json:"template-scaleout,omitempty"`
	TemplateServerSSL                string                 `json:"template-server-ssl,omitempty"`
	TemplateServerSSLShared          string                 `json:"template-server-ssl-shared,omitempty"`
	TemplateSip                      string                 `json:"template-sip,omitempty"`
	TemplateSmpp                     string                 `json:"template-smpp,omitempty"`
	TemplateSMTP                     string                 `json:"template-smtp,omitempty"`
	TemplateSsli                     string                 `json:"template-ssli,omitempty"`
	TemplateTCP                      string                 `json:"template-tcp,omitempty"`
	TemplateTCPProxy                 string                 `json:"template-tcp-proxy,omitempty"`
	TemplateTCPProxyClient           string                 `json:"template-tcp-proxy-client,omitempty"`
	TemplateTCPProxyServer           string                 `json:"template-tcp-proxy-server,omitempty"`
	TemplateTCPShared                string                 `json:"template-tcp-shared,omitempty"`
	TemplateUDP                      string                 `json:"template-udp,omitempty"`
	TemplateUDPShared                string                 `json:"template-udp-shared,omitempty"`
	TemplateVirtualPort              string                 `json:"template-virtual-port,omitempty"`
	TrunkFwd                         int                    `json:"trunk-fwd,omitempty"`
	TrunkRev                         int                    `json:"trunk-rev,omitempty"`
	UseAlternatePort                 shared.Boolean         `json:"use-alternate-port,omitempty"`
	UseCgnv6                         shared.Boolean         `json:"use-cgnv6,omitempty"`
	UseDefaultIfNoServer             shared.Boolean         `json:"use-default-if-no-server,omitempty"`
	UseRcvHopForResp                 shared.Boolean         `json:"use-rcv-hop-for-resp,omitempty"`
	View                             int                    `json:"view,omitempty"`
	WafTemplate                      string                 `json:"waf-template,omitempty"`
	WhenDown                         shared.Boolean         `json:"when-down,omitempty"`
	WhenDownProtocol2                shared.Boolean         `json:"when-down-protocol2,omitempty"`
}
type ListObjects []Object

// ACLNameList Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#acl-name-list
type ACLNameList []struct {
	ACLName           string `json:"acl-name,omitempty"`
	ACLNameSeqNum     int    `json:"acl-name-seq-num,omitempty"`
	ACLNameSrcNatPool string `json:"acl-name-src-nat-pool,omitempty"`
}

// ACLIDList Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#acl-id-list
type ACLIDList []struct {
	ACLID           int    `json:"acl-id,omitempty"`
	ACLIDSeqNum     int    `json:"acl-id-seq-num,omitempty"`
	ACLIDSrcNatPool string `json:"acl-id-src-nat-pool,omitempty"`
}

// AflexScripts Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#aflex-scripts
type AflexScripts []struct {
	Aflex string `json:"aflex,omitempty"`
}

// AuthCfg Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#auth-cfg
type AuthCfg struct {
	AAAPolicy string `json:"aaa-policy,omitempty"`
}
