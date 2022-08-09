package virtualserver

import "github.com/ureuzy/acos-client-go/pkg/axapi/slb/virtualserverport"

type ListBody struct {
	ListObjects `json:"virtual-server-list"`
}

type Body struct {
	Object `json:"virtual-server"`
}

// Object Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server.html#virtual-server-attributes
type Object struct {
	AclID                 int                            `json:"acl-id,omitempty"`
	AclName               string                         `json:"acl-name,omitempty"`
	ArpDisable            int                            `json:"arp-disable,omitempty"`
	Description           string                         `json:"description,omitempty"`
	DisableVipAdv         int                            `json:"disable-vip-adv,omitempty"`
	EnableDisableAction   string                         `json:"enable-disable-action,omitempty"`
	Ethernet              int                            `json:"ethernet,omitempty"`
	ExtendedStats         int                            `json:"extended-stats,omitempty"`
	HADynamic             int                            `json:"ha-dynamic,omitempty"`
	IPAddress             string                         `json:"ip-address"`
	IPv6Acl               string                         `json:"ipv6-acl,omitempty"`
	IPv6Address           string                         `json:"ipv6-address,omitempty"`
	MigrateVIP            *MigrateVip                    `json:"migrate-vip,omitempty"`
	Name                  string                         `json:"name"`
	Netmask               string                         `json:"netmask,omitempty"`
	PortList              *virtualserverport.ListObjects `json:"port-list,omitempty"`
	RedistributeRouteMap  string                         `json:"redistribute-route-map,omitempty"`
	RedistributionFlagged int                            `json:"redistribution-flagged,omitempty"`
	StatsDataAction       string                         `json:"stats-data-action,omitempty"`
	TemplateLogging       string                         `json:"template-logging,omitempty"`
	TemplatePolicy        string                         `json:"template-policy,omitempty"`
	TemplateScaleout      string                         `json:"template-scaleout,omitempty"`
	TemplateVirtualServer string                         `json:"template-virtual-server,omitempty"`
	UseIfIP               int                            `json:"use-if-ip,omitempty"`
	UserTag               string                         `json:"user-tag,omitempty"`
	UUID                  string                         `json:"uuid,omitempty"`
	VRID                  int                            `json:"vrid,omitempty"`
}
type ListObjects []Object

// MigrateVip Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server.html#migrate-vip
type MigrateVip struct {
	CancelMigration    bool   `json:"cancel-migration,omitempty"`
	FinishMigration    bool   `json:"finish-migration,omitempty"`
	TargetDataCpu      int    `json:"target-data-cpu,omitempty"`
	TargetFloatingIPv4 string `json:"target-floating-ipv4,omitempty"`
	TargetFloatingIPv6 string `json:"target-floating-ipv6,omitempty"`
	UUID               string `json:"uuid,omitempty"`
}
