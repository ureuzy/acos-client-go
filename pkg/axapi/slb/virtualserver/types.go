package virtualserver

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/shared"
	"github.com/ureuzy/acos-client-go/pkg/axapi/slb/virtualserver/port"
)

// Object Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server.html#virtual-server-attributes
type Object struct {
	shared.AxaBase        `json:",inline"`
	ACLID                 int               `json:"acl-id,omitempty"`
	ACLName               string            `json:"acl-name,omitempty"`
	ArpDisable            shared.Boolean    `json:"arp-disable,omitempty"`
	Description           string            `json:"description,omitempty"`
	DisableVipAdv         shared.Boolean    `json:"disable-vip-adv,omitempty"`
	EnableDisableAction   string            `json:"enable-disable-action,omitempty"`
	Ethernet              int               `json:"ethernet,omitempty"`
	ExtendedStats         shared.Boolean    `json:"extended-stats,omitempty"`
	HADynamic             int               `json:"ha-dynamic,omitempty"`
	IPAddress             string            `json:"ip-address"`
	IPv6ACL               string            `json:"ipv6-acl,omitempty"`
	IPv6Address           string            `json:"ipv6-address,omitempty"`
	MigrateVIP            *MigrateVip       `json:"migrate-vip,omitempty"`
	Name                  string            `json:"name"`
	Netmask               string            `json:"netmask,omitempty"`
	PortList              *port.ListObjects `json:"port-list,omitempty"`
	RedistributeRouteMap  string            `json:"redistribute-route-map,omitempty"`
	RedistributionFlagged shared.Boolean    `json:"redistribution-flagged,omitempty"`
	StatsDataAction       string            `json:"stats-data-action,omitempty"`
	TemplateLogging       string            `json:"template-logging,omitempty"`
	TemplatePolicy        string            `json:"template-policy,omitempty"`
	TemplateScaleout      string            `json:"template-scaleout,omitempty"`
	TemplateVirtualServer string            `json:"template-virtual-server,omitempty"`
	UseIfIP               shared.Boolean    `json:"use-if-ip,omitempty"`
	VRID                  int               `json:"vrid,omitempty"`
}
type ListObjects []Object

// MigrateVip Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/slb_virtual_server.html#migrate-vip
type MigrateVip struct {
	shared.AxaBase     `json:",inline"`
	CancelMigration    shared.Boolean `json:"cancel-migration,omitempty"`
	FinishMigration    shared.Boolean `json:"finish-migration,omitempty"`
	TargetDataCPU      int            `json:"target-data-cpu,omitempty"`
	TargetFloatingIPv4 string         `json:"target-floating-ipv4,omitempty"`
	TargetFloatingIPv6 string         `json:"target-floating-ipv6,omitempty"`
}
