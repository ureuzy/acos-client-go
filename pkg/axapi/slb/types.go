package slb

type VirtualServerList []VirtualServer

// VirtualServer Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server.html#virtual-server-attributes
type VirtualServer struct {
	AclID                 int         `json:"acl-id,omitempty"`
	AclName               string      `json:"acl-name,omitempty"`
	ArpDisable            int         `json:"arp-disable,omitempty"`
	Description           string      `json:"description,omitempty"`
	DisableVipAdv         int         `json:"disable-vip-adv,omitempty"`
	EnableDisableAction   string      `json:"enable-disable-action,omitempty"`
	Ethernet              int         `json:"ethernet,omitempty"`
	ExtendedStats         int         `json:"extended-stats,omitempty"`
	HADynamic             int         `json:"ha-dynamic,omitempty"`
	IPAddress             string      `json:"ip-address"`
	IPv6Acl               string      `json:"ipv6-acl,omitempty"`
	IPv6Address           string      `json:"ipv6-address,omitempty"`
	MigrateVIP            *MigrateVip `json:"migrate-vip,omitempty"`
	Name                  string      `json:"name"`
	Netmask               string      `json:"netmask,omitempty"`
	PortList              *PortList   `json:"port-list,omitempty"`
	RedistributeRouteMap  string      `json:"redistribute-route-map,omitempty"`
	RedistributionFlagged int         `json:"redistribution-flagged,omitempty"`
	StatsDataAction       string      `json:"stats-data-action,omitempty"`
	TemplateLogging       string      `json:"template-logging,omitempty"`
	TemplatePolicy        string      `json:"template-policy,omitempty"`
	TemplateScaleout      string      `json:"template-scaleout,omitempty"`
	TemplateVirtualServer string      `json:"template-virtual-server,omitempty"`
	UseIfIP               int         `json:"use-if-ip,omitempty"`
	UserTag               string      `json:"user-tag,omitempty"`
	UUID                  string      `json:"uuid,omitempty"`
	VRID                  int         `json:"vrid,omitempty"`
}

// MigrateVip Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server.html#migrate-vip
type MigrateVip struct {
	CancelMigration    bool   `json:"cancel-migration"`
	FinishMigration    bool   `json:"finish-migration"`
	TargetDataCpu      int    `json:"target-data-cpu"`
	TargetFloatingIPv4 string `json:"target-floating-ipv4"`
	TargetFloatingIPv6 string `json:"target-floating-ipv6"`
	UUID               string `json:"uuid"`
}

type PortList []Port

// Port Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#port-attributes
type Port struct {
	AclIDList                        *AclIDList          `json:"acl-id-list"`
	AclNameList                      *AclNameList        `json:"acl-name-list"`
	Action                           string              `json:"action"`
	AflexScripts                     *AflexScripts       `json:"aflex-scripts"`
	AltProtocol1                     string              `json:"alt-protocol1"`
	AltProtocol2                     string              `json:"alt-protocol2"`
	AlternatePort                    int                 `json:"alternate-port"`
	AlternatePortNumber              int                 `json:"alternate-port-number"`
	AuthCfg                          *AuthCfg            `json:"auth-cfg"`
	Auto                             int                 `json:"auto"`
	ClientIPStickyNat                int                 `json:"clientip-sticky-nat"`
	ConnLimit                        int                 `json:"conn-limit"`
	DefSelectionIfPrefFailed         string              `json:"def-selection-if-pref-failed"`
	EnablePlayerIDCheck              int                 `json:"enable-playerid-check"`
	EthFwd                           int                 `json:"eth-fwd"`
	EthRev                           int                 `json:"eth-rev"`
	Expand                           int                 `json:"expand"`
	ExtendedStats                    int                 `json:"extended-stats"`
	ForceRoutingMode                 int                 `json:"force-routing-mode"`
	GslbEnable                       int                 `json:"gslb-enable"`
	HAConnMirror                     int                 `json:"ha-conn-mirror"`
	IPInIP                           int                 `json:"ipinip"`
	L7HardwareAssist                 int                 `json:"l7-hardware-assist"`
	MessageSwitching                 int                 `json:"message-switching"`
	Name                             string              `json:"name"`
	NoAutoUpOnAflex                  int                 `json:"no-auto-up-on-aflex"`
	NoDestNat                        int                 `json:"no-dest-nat"`
	NoLogging                        int                 `json:"no-logging"`
	OnSyn                            int                 `json:"on-syn"`
	PersistType                      string              `json:"persist-type"`
	Pool                             string              `json:"pool"`
	PortNumber                       int                 `json:"port-number"`
	PortTranslation                  int                 `json:"port-translation"`
	Precedence                       int                 `json:"precedence"`
	Protocol                         string              `json:"protocol"`
	Range                            int                 `json:"range"`
	Rate                             int                 `json:"rate"`
	RedirectToHttps                  int                 `json:"redirect-to-https"`
	ReqFail                          int                 `json:"req-fail"`
	Reset                            int                 `json:"reset"`
	ResetOnServerSelectionFail       int                 `json:"reset-on-server-selection-fail"`
	RtpSipCallIDMatch                int                 `json:"rtp-sip-call-id-match"`
	SamplingEnable                   *SamplingEnableList `json:"sampling-enable"`
	ScaleOutBucketCount              int                 `json:"scaleout-bucket-count"`
	ScaleOutDeviceGroup              int                 `json:"scaleout-device-group"`
	Secs                             int                 `json:"secs"`
	ServSelFail                      int                 `json:"serv-sel-fail"`
	ServiceGroup                     string              `json:"service-group"`
	SharedPartitionClientSSLTemplate int                 `json:"shared-partition-client-ssl-template"`
	SharedPartitionHTTPTemplate      int                 `json:"shared-partition-http-template"`
	SharedPartitionServerSSLTemplate int                 `json:"shared-partition-server-ssl-template"`
	SharedPartitionTCP               int                 `json:"shared-partition-tcp"`
	SharedPartitionUDP               int                 `json:"shared-partition-udp"`
	SkipRevHash                      int                 `json:"skip-rev-hash"`
	SNatOnVip                        int                 `json:"snat-on-vip"`
	StatsDataAction                  string              `json:"stats-data-action"`
	SupportHttp2                     int                 `json:"support-http2"`
	SynCookie                        int                 `json:"syn-cookie"`
	TemplateCache                    string              `json:"template-cache"`
	TemplateClientSSL                string              `json:"template-client-ssl"`
	TemplateClientSSLShared          string              `json:"template-client-ssl-shared"`
	TemplateConnectionReuse          string              `json:"template-connection-reuse"`
	TemplateDBLB                     string              `json:"template-dblb"`
	TemplateDiameter                 string              `json:"template-diameter"`
	TemplateDNS                      string              `json:"template-dns"`
	TemplateDynamicService           string              `json:"template-dynamic-service"`
	TemplateExternalService          string              `json:"template-external-service"`
	TemplateFileInspection           string              `json:"template-file-inspection"`
	TemplateFix                      string              `json:"template-fix"`
	TemplateFTP                      string              `json:"template-ftp"`
	TemplateHTTP                     string              `json:"template-http"`
	TemplateHTTPPolicy               string              `json:"template-http-policy"`
	TemplateHTTPShared               string              `json:"template-http-shared"`
	TemplateImapPop3                 string              `json:"template-imap-pop3"`
	TemplatePersistCookie            string              `json:"template-persist-cookie"`
	TemplatePersistDestinationIP     string              `json:"template-persist-destination-ip"`
	TemplatePersistSourceIP          string              `json:"template-persist-source-ip"`
	TemplatePersistSSLSid            string              `json:"template-persist-ssl-sid"`
	TemplatePolicy                   string              `json:"template-policy"`
	TemplateReqmodIcap               string              `json:"template-reqmod-icap"`
	TemplateRespmodIcap              string              `json:"template-respmod-icap"`
	TemplateScaleout                 string              `json:"template-scaleout"`
	TemplateServerSSL                string              `json:"template-server-ssl"`
	TemplateServerSSLShared          string              `json:"template-server-ssl-shared"`
	TemplateSip                      string              `json:"template-sip"`
	TemplateSmpp                     string              `json:"template-smpp"`
	TemplateSmtp                     string              `json:"template-smtp"`
	TemplateSsli                     string              `json:"template-ssli"`
	TemplateTCP                      string              `json:"template-tcp"`
	TemplateTCPProxy                 string              `json:"template-tcp-proxy"`
	TemplateTCPProxyClient           string              `json:"template-tcp-proxy-client"`
	TemplateTCPProxyServer           string              `json:"template-tcp-proxy-server"`
	TemplateTCPShared                string              `json:"template-tcp-shared"`
	TemplateUDP                      string              `json:"template-udp"`
	TemplateUDPShared                string              `json:"template-udp-shared"`
	TemplateVirtualPort              string              `json:"template-virtual-port"`
	TrunkFwd                         int                 `json:"trunk-fwd"`
	TrunkRev                         int                 `json:"trunk-rev"`
	UseAlternatePort                 int                 `json:"use-alternate-port"`
	UseCgnv6                         int                 `json:"use-cgnv6"`
	UseDefaultIfNoServer             int                 `json:"use-default-if-no-server"`
	UseRcvHopForResp                 int                 `json:"use-rcv-hop-for-resp"`
	UserTag                          string              `json:"user-tag"`
	UUID                             string              `json:"uuid"`
	View                             int                 `json:"view"`
	WafTemplate                      string              `json:"waf-template"`
	WhenDown                         int                 `json:"when-down"`
	WhenDownProtocol2                int                 `json:"when-down-protocol2"`
}

// AclName Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#acl-name-list
type AclName struct {
	AclName           string `json:"acl-name"`
	AclNameSeqNum     int    `json:"acl-name-seq-num"`
	AclNameSrcNatPool string `json:"acl-name-src-nat-pool"`
}
type AclNameList []AclName

// SamplingEnable Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#sampling-enable
type SamplingEnable struct {
	Counters1 string `json:"counters1"`
}
type SamplingEnableList []SamplingEnable

// AclID Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#acl-id-list
type AclID struct {
	AclID           int    `json:"acl-id"`
	AclIDSeqNum     int    `json:"acl-id-seq-num"`
	AclIDSrcNatPool string `json:"acl-id-src-nat-pool"`
}
type AclIDList []AclID

// AflexScript Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#aflex-scripts
type AflexScript struct {
	Aflex string `json:"aflex"`
}
type AflexScripts []AflexScript

// AuthCfg Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server_port.html#auth-cfg
type AuthCfg struct {
	AAAPolicy string `json:"aaa-policy"`
}
