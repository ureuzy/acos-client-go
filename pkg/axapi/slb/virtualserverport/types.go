package virtualserverport

type ListObject struct {
	VirtualServerPortList `json:"port-list"`
}

type Object struct {
	VirtualServerPort `json:"port"`
}

// VirtualServerPort Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#port-attributes
type VirtualServerPort struct {
	AclIDList                        *AclIDList          `json:"acl-id-list,omitempty"`
	AclNameList                      *AclNameList        `json:"acl-name-list,omitempty"`
	Action                           string              `json:"action,omitempty"`
	AflexScripts                     *AflexScripts       `json:"aflex-scripts,omitempty"`
	AltProtocol1                     string              `json:"alt-protocol1,omitempty"`
	AltProtocol2                     string              `json:"alt-protocol2,omitempty"`
	AlternatePort                    int                 `json:"alternate-port,omitempty"`
	AlternatePortNumber              int                 `json:"alternate-port-number,omitempty"`
	AuthCfg                          *AuthCfg            `json:"auth-cfg,omitempty"`
	Auto                             int                 `json:"auto,omitempty"`
	ClientIPStickyNat                int                 `json:"clientip-sticky-nat,omitempty"`
	ConnLimit                        int                 `json:"conn-limit,omitempty"`
	DefSelectionIfPrefFailed         string              `json:"def-selection-if-pref-failed,omitempty"`
	EnablePlayerIDCheck              int                 `json:"enable-playerid-check,omitempty"`
	EthFwd                           int                 `json:"eth-fwd,omitempty"`
	EthRev                           int                 `json:"eth-rev,omitempty"`
	Expand                           int                 `json:"expand,omitempty"`
	ExtendedStats                    int                 `json:"extended-stats,omitempty"`
	ForceRoutingMode                 int                 `json:"force-routing-mode,omitempty"`
	GslbEnable                       int                 `json:"gslb-enable,omitempty"`
	HAConnMirror                     int                 `json:"ha-conn-mirror,omitempty"`
	IPInIP                           int                 `json:"ipinip,omitempty"`
	L7HardwareAssist                 int                 `json:"l7-hardware-assist,omitempty"`
	MessageSwitching                 int                 `json:"message-switching,omitempty"`
	Name                             string              `json:"name,omitempty"`
	NoAutoUpOnAflex                  int                 `json:"no-auto-up-on-aflex,omitempty"`
	NoDestNat                        int                 `json:"no-dest-nat,omitempty"`
	NoLogging                        int                 `json:"no-logging,omitempty"`
	OnSyn                            int                 `json:"on-syn,omitempty"`
	PersistType                      string              `json:"persist-type,omitempty"`
	Pool                             string              `json:"pool,omitempty"`
	PortNumber                       int                 `json:"port-number,omitempty"`
	PortTranslation                  int                 `json:"port-translation,omitempty"`
	Precedence                       int                 `json:"precedence,omitempty"`
	Protocol                         string              `json:"protocol,omitempty"`
	Range                            int                 `json:"range,omitempty"`
	Rate                             int                 `json:"rate,omitempty"`
	RedirectToHttps                  int                 `json:"redirect-to-https,omitempty"`
	ReqFail                          int                 `json:"req-fail,omitempty"`
	Reset                            int                 `json:"reset,omitempty"`
	ResetOnServerSelectionFail       int                 `json:"reset-on-server-selection-fail,omitempty"`
	RtpSipCallIDMatch                int                 `json:"rtp-sip-call-id-match,omitempty"`
	SamplingEnable                   *SamplingEnableList `json:"sampling-enable,omitempty"`
	ScaleOutBucketCount              int                 `json:"scaleout-bucket-count,omitempty"`
	ScaleOutDeviceGroup              int                 `json:"scaleout-device-group,omitempty"`
	Secs                             int                 `json:"secs,omitempty"`
	ServSelFail                      int                 `json:"serv-sel-fail,omitempty"`
	ServiceGroup                     string              `json:"service-group,omitempty"`
	SharedPartitionClientSSLTemplate int                 `json:"shared-partition-client-ssl-template,omitempty"`
	SharedPartitionHTTPTemplate      int                 `json:"shared-partition-http-template,omitempty"`
	SharedPartitionServerSSLTemplate int                 `json:"shared-partition-server-ssl-template,omitempty"`
	SharedPartitionTCP               int                 `json:"shared-partition-tcp,omitempty"`
	SharedPartitionUDP               int                 `json:"shared-partition-udp,omitempty"`
	SkipRevHash                      int                 `json:"skip-rev-hash,omitempty"`
	SNatOnVip                        int                 `json:"snat-on-vip,omitempty"`
	StatsDataAction                  string              `json:"stats-data-action,omitempty"`
	SupportHttp2                     int                 `json:"support-http2,omitempty"`
	SynCookie                        int                 `json:"syn-cookie,omitempty"`
	TemplateCache                    string              `json:"template-cache,omitempty"`
	TemplateClientSSL                string              `json:"template-client-ssl,omitempty"`
	TemplateClientSSLShared          string              `json:"template-client-ssl-shared,omitempty"`
	TemplateConnectionReuse          string              `json:"template-connection-reuse,omitempty"`
	TemplateDBLB                     string              `json:"template-dblb,omitempty"`
	TemplateDiameter                 string              `json:"template-diameter,omitempty"`
	TemplateDNS                      string              `json:"template-dns,omitempty"`
	TemplateDynamicService           string              `json:"template-dynamic-service,omitempty"`
	TemplateExternalService          string              `json:"template-external-service,omitempty"`
	TemplateFileInspection           string              `json:"template-file-inspection,omitempty"`
	TemplateFix                      string              `json:"template-fix,omitempty"`
	TemplateFTP                      string              `json:"template-ftp,omitempty"`
	TemplateHTTP                     string              `json:"template-http,omitempty"`
	TemplateHTTPPolicy               string              `json:"template-http-policy,omitempty"`
	TemplateHTTPShared               string              `json:"template-http-shared,omitempty"`
	TemplateImapPop3                 string              `json:"template-imap-pop3,omitempty"`
	TemplatePersistCookie            string              `json:"template-persist-cookie,omitempty"`
	TemplatePersistDestinationIP     string              `json:"template-persist-destination-ip,omitempty"`
	TemplatePersistSourceIP          string              `json:"template-persist-source-ip,omitempty"`
	TemplatePersistSSLSid            string              `json:"template-persist-ssl-sid,omitempty"`
	TemplatePolicy                   string              `json:"template-policy,omitempty"`
	TemplateReqmodIcap               string              `json:"template-reqmod-icap,omitempty"`
	TemplateRespmodIcap              string              `json:"template-respmod-icap,omitempty"`
	TemplateScaleout                 string              `json:"template-scaleout,omitempty"`
	TemplateServerSSL                string              `json:"template-server-ssl,omitempty"`
	TemplateServerSSLShared          string              `json:"template-server-ssl-shared,omitempty"`
	TemplateSip                      string              `json:"template-sip,omitempty"`
	TemplateSmpp                     string              `json:"template-smpp,omitempty"`
	TemplateSmtp                     string              `json:"template-smtp,omitempty"`
	TemplateSsli                     string              `json:"template-ssli,omitempty"`
	TemplateTCP                      string              `json:"template-tcp,omitempty"`
	TemplateTCPProxy                 string              `json:"template-tcp-proxy,omitempty"`
	TemplateTCPProxyClient           string              `json:"template-tcp-proxy-client,omitempty"`
	TemplateTCPProxyServer           string              `json:"template-tcp-proxy-server,omitempty"`
	TemplateTCPShared                string              `json:"template-tcp-shared,omitempty"`
	TemplateUDP                      string              `json:"template-udp,omitempty"`
	TemplateUDPShared                string              `json:"template-udp-shared,omitempty"`
	TemplateVirtualPort              string              `json:"template-virtual-port,omitempty"`
	TrunkFwd                         int                 `json:"trunk-fwd,omitempty"`
	TrunkRev                         int                 `json:"trunk-rev,omitempty"`
	UseAlternatePort                 int                 `json:"use-alternate-port,omitempty"`
	UseCgnv6                         int                 `json:"use-cgnv6,omitempty"`
	UseDefaultIfNoServer             int                 `json:"use-default-if-no-server,omitempty"`
	UseRcvHopForResp                 int                 `json:"use-rcv-hop-for-resp,omitempty"`
	UserTag                          string              `json:"user-tag,omitempty"`
	UUID                             string              `json:"uuid,omitempty"`
	View                             int                 `json:"view,omitempty"`
	WafTemplate                      string              `json:"waf-template,omitempty"`
	WhenDown                         int                 `json:"when-down,omitempty"`
	WhenDownProtocol2                int                 `json:"when-down-protocol2,omitempty"`
}
type VirtualServerPortList []VirtualServerPort

// AclName Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#acl-name-list
type AclName struct {
	AclName           string `json:"acl-name,omitempty"`
	AclNameSeqNum     int    `json:"acl-name-seq-num,omitempty"`
	AclNameSrcNatPool string `json:"acl-name-src-nat-pool,omitempty"`
}
type AclNameList []AclName

// SamplingEnable Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#sampling-enable
type SamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SamplingEnableList []SamplingEnable

// AclID Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#acl-id-list
type AclID struct {
	AclID           int    `json:"acl-id,omitempty"`
	AclIDSeqNum     int    `json:"acl-id-seq-num,omitempty"`
	AclIDSrcNatPool string `json:"acl-id-src-nat-pool,omitempty"`
}
type AclIDList []AclID

// AflexScript Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#aflex-scripts
type AflexScript struct {
	Aflex string `json:"aflex,omitempty"`
}
type AflexScripts []AflexScript

// AuthCfg Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#auth-cfg
type AuthCfg struct {
	AAAPolicy string `json:"aaa-policy,omitempty"`
}
