package ipserver

import "github.com/ureuzy/acos-client-go/pkg/axapi/shared"

type IPServer struct {
	shared.AxaBase `json:",inline"`
	IPServerName   string `json:"ip-server-name,omitempty"`
}
