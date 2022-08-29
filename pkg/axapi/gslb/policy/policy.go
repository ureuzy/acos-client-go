package policy

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/axapi/shared"
	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_policy.html

func New(c utils.HTTPClient, basePath string) rest.Operator[Body, ListBody] {
	const path = "policy"
	return rest.Rest[Body, ListBody](c, fmt.Sprintf("%s/%s", basePath, path))
}

type ListBody struct {
	ListObjects `json:"policy-list"`
}

type Body struct {
	Object `json:"policy"`
}

type ListObjects []Object

// Object Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/gslb_policy.html#policy-attributes

type Object struct {
	shared.AxaBase              `json:",inline"`
	Name                        string           `json:"name,omitempty"`
	HealthCheck                 bool             `json:"health-check,omitempty"`
	HealthCheckPreferenceEnable bool             `json:"health-check-preference-enable,omitempty"`
	HealthPreferenceTop         int              `json:"health-preference-top,omitempty"`
	AmountFirst                 bool             `json:"amount-first,omitempty"`
	WeightedIPEnable            bool             `json:"weighted-ip-enable,omitempty"`
	WeightedIPTotalHits         bool             `json:"weighted-ip-total-hits,omitempty"`
	WeightedSiteEnable          bool             `json:"weighted-site-enable,omitempty"`
	WeightedSiteTotalHits       bool             `json:"weighted-site-total-hits,omitempty"`
	WeightedAlias               bool             `json:"weighted-alias,omitempty"`
	ActiveServersEnable         bool             `json:"active-servers-enable,omitempty"`
	ActiveServersFailBreak      bool             `json:"active-servers-fail-break,omitempty"`
	BwCostEnable                bool             `json:"bw-cost-enable,omitempty"`
	BwCostFailBreak             bool             `json:"bw-cost-fail-break,omitempty"`
	Geographic                  bool             `json:"geographic,omitempty"`
	NumSessionEnable            bool             `json:"num-session-enable,omitempty"`
	NumSessionTolerance         int              `json:"num-session-tolerance,omitempty"`
	AdminPreference             bool             `json:"admin-preference,omitempty"`
	AliasAdminPreference        bool             `json:"alias-admin-preference,omitempty"`
	LeastResponse               bool             `json:"least-response,omitempty"`
	AdminIPEnable               bool             `json:"admin-ip-enable,omitempty"`
	AdminIPTopOnly              bool             `json:"admin-ip-top-only,omitempty"`
	OrderedIPTopOnly            bool             `json:"ordered-ip-top-only,omitempty"`
	RoundRobin                  bool             `json:"round-robin,omitempty"`
	MetricForceCheck            bool             `json:"metric-force-check,omitempty"`
	IPList                      string           `json:"ip-list,omitempty"`
	MetricOrder                 bool             `json:"metric-order,omitempty"`
	MetricType                  bool             `json:"metric-type,omitempty"` // "enum":[  "health-check",  "weighted-ip",  "weighted-site",  "capacity",  "active-servers",  "active-rdt",  "geographic",  "connection-load",  "num-session",  "admin-preference",  "bw-cost",  "least-response",  "admin-ip"]
	Capacity                    Capacity         `json:"capacity,omitempty"`
	ConnectionLoad              ConnectionLoad   `json:"connection-load,omitempty"`
	DNS                         DNS              `json:"dns,omitempty"`
	GeoLocationList             []GeoLocation    `json:"geo-location-list,omitempty"`
	GeoLocationMatch            GeoLocationMatch `json:"geo-location-match,omitempty"`
	ActiveRdt                   ActiveRdt        `json:"active-rdt,omitempty"`
	AutoMap                     AutoMap          `json:"auto-map,omitempty"`
	EDNS                        EDNS             `json:"edns,omitempty"`
}

type Capacity struct {
	shared.AxaBase `json:",inline"`
	// TODO
}

type ConnectionLoad struct {
	shared.AxaBase `json:",inline"`
	// TODO
}

type DNS struct {
	shared.AxaBase `json:",inline"`
	// TODO
}

type GeoLocation struct {
	shared.AxaBase `json:",inline"`
	// TODO
}

type GeoLocationMatch struct {
	shared.AxaBase `json:",inline"`
	// TODO
}

type ActiveRdt struct {
	shared.AxaBase `json:",inline"`
	// TODO
}

type AutoMap struct {
	shared.AxaBase `json:",inline"`
	// TODO
}

type EDNS struct {
	shared.AxaBase `json:",inline"`
	// TODO
}
