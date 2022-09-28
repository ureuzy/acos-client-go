package site

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/site/ipserver"
	"github.com/ureuzy/acos-client-go/pkg/axapi/shared"
)

// Object Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_site.html#site-attributes
type Object struct {
	shared.AxaBase       `json:",inline"`
	SiteName             string              `json:"site-name,omitempty"`
	AutoMap              shared.Boolean      `json:"auto-map,omitempty"`
	Disable              shared.Boolean      `json:"disable,omitempty"`
	Weight               int                 `json:"weight,omitempty"`
	MultipleGeoLocations []string            `json:"multiple-geo-locations,omitempty"`
	Template             string              `json:"template,omitempty"`
	BwCost               int                 `json:"bw-cost,omitempty"`
	Limit                int                 `json:"limit,omitempty"`
	Threshold            int                 `json:"threshold,omitempty"`
	ProtoAgingTime       int                 `json:"proto-aging-time,omitempty"`
	ProtoAgingFast       shared.Boolean      `json:"proto-aging-fast,omitempty"`
	Controller           string              `json:"controller,omitempty"`
	IPServerList         []ipserver.IPServer `json:"ip-server-list,omitempty"`
	ActiveRdt            ActiveRdt           `json:"active-rdt,omitempty"`
	EasyRdt              EasyRdt             `json:"easy-rdt,omitempty"`
	SlbDevList           []SlbDevList        `json:"slb-dev-list,omitempty"`
}

type ActiveRdt struct {
	shared.AxaBase `json:",inline"`
	// TODO
}

type EasyRdt struct {
	shared.AxaBase `json:",inline"`
	// TODO
}

type SlbDevList struct {
	shared.AxaBase `json:",inline"`
	// TODO
}
