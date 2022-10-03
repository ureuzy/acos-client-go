package gslb

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/policy"
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/serviceip"
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/site"
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/site/siteipserver"
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/zone"
	"github.com/ureuzy/acos-client-go/pkg/axapi/gslb/zone/zoneservice"
	"github.com/ureuzy/acos-client-go/pkg/rest"
	"github.com/ureuzy/acos-client-go/utils"
)

const path = "gslb"

type Operator struct {
	Policy       rest.Operator[policy.Body, policy.ListBody]
	ServiceIP    rest.Operator[serviceip.Body, serviceip.ListBody]
	Site         rest.Operator[site.Body, site.ListBody]
	SiteIPServer rest.Operator[siteipserver.Body, siteipserver.ListBody]
	Zone         rest.Operator[zone.Body, zone.ListBody]
	ZoneService  rest.Operator[zoneservice.Body, zoneservice.ListBody]
}

func New(c utils.HTTPClient) *Operator {
	return &Operator{
		Policy:       policy.New(c, path),
		ServiceIP:    serviceip.New(c, path),
		Site:         site.New(c, path),
		SiteIPServer: siteipserver.New(c, path),
		Zone:         zone.New(c, path),
		ZoneService:  zoneservice.New(c, path),
	}
}
