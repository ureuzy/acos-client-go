package healthmonitor

import "github.com/ureuzy/acos-client-go/pkg/axapi/shared"

type Object struct {
	shared.AxaBase             `json:",inline"`
	Name                       string         `json:"name,omitempty"`
	DsrL2Strict                shared.Boolean `json:"dsr-l2-strict,omitempty"`
	Retry                      int            `json:"retry,omitempty"`
	UpRetry                    int            `json:"up-retry,omitempty"`
	OverrideIPv4               string         `json:"override-ipv4,omitempty"`
	OverrideIPv6               string         `json:"override-ipv6,omitempty"`
	OverridePort               int            `json:"override-port,omitempty"`
	Passive                    shared.Boolean `json:"passive,omitempty"`
	StatusCode                 string         `json:"status-code,omitempty"` //"enum":["status-code-2xx","status-code-non-5xx"]
	PassiveInterval            int            `json:"passive-interval,omitempty"`
	SampleThreshold            int            `json:"sample-threshold,omitempty"`
	Threshold                  int            `json:"threshold,omitempty"`
	StrictRetryOnServerErrResp shared.Boolean `json:"strict-retry-on-server-err-resp,omitempty"`
	DisableAfterDown           shared.Boolean `json:"disable-after-down,omitempty"`
	Interval                   int            `json:"interval,omitempty"`
	Timeout                    int            `json:"timeout,omitempty"`
	SslCiphers                 string         `json:"ssl-ciphers,omitempty"`
	SslTicket                  shared.Boolean `json:"ssl-ticket,omitempty"`
	SslTicketLifetime          int            `json:"ssl-ticket-lifetime,omitempty"`
	SslVersion                 int            `json:"ssl-version,omitempty"`
	SslDgVersion               int            `json:"ssl-dgversion,omitempty"`
	Method                     Method         `json:"method,omitempty"`
	A10URL                     string         `json:"a10-url,omitempty"`
}

type HTTPMethod string

type Method struct {
	TCP   *TCP   `json:"tcp,omitempty"`
	HTTP  *HTTP  `json:"http,omitempty"`
	HTTPS *HTTPS `json:"https,omitempty"`
	ICMP  *ICMP  `json:"icmp,omitempty"`
}

type HTTP struct {
	shared.AxaBase        `json:",inline"`
	HTTP                  shared.Boolean `json:"http,omitempty"`
	HTTPPort              int            `json:"http-port,omitempty"`
	HTTPExpect            shared.Boolean `json:"http-expect,omitempty"`
	HTTPResponseCode      string         `json:"http-response-code,omitempty"`
	HTTPResponseCodeRegex string         `json:"response-code-regex,omitempty"`
	HTTPText              string         `json:"http-text,omitempty"`
	TextRegex             string         `json:"text-regex,omitempty"`
	HTTPHost              string         `json:"http-host,omitempty"`
	HTTPMaintenanceCode   string         `json:"http-maintenance-code,omitempty"`
	HTTPURL               shared.Boolean `json:"http-url,omitempty"`
	URLType               HTTPMethod     `json:"url-type,omitempty"`
	Maintenance           shared.Boolean `json:"maintenance,omitempty"`
	MaintenanceText       string         `json:"maintenance-text,omitempty"`
	MaintenanceTextRegex  string         `json:"maintenance-text-regex,omitempty"`
	URLPath               string         `json:"url-path,omitempty"`
	PostPath              string         `json:"post-path,omitempty"`
	PostType              string         `json:"post-type,omitempty"`
	HTTPPostData          string         `json:"http-postdata,omitempty"`
	HTTPPostfile          string         `json:"http-postfile,omitempty"`
	HTTPUsername          string         `json:"http-username,omitempty"`
	HTTPPassword          string         `json:"http-password,omitempty"`
	HTTPPasswordString    string         `json:"http-password-string,omitempty"`
	HTTPKerberosAuth      shared.Boolean `json:"http-kerberos-auth,omitempty"`
	HTTPKerberosRealm     string         `json:"http-kerberos-realm,omitempty"`
	HTTPKerberosKdc       KerberosKdc    `json:"http-kerberos-kdc,omitempty"`
}

type HTTPS struct {
	shared.AxaBase         `json:",inline"`
	HTTP                   shared.Boolean `json:"https,omitempty"`
	WebPort                int            `json:"web-port,omitempty"`
	DisableSslV2Hello      shared.Boolean `json:"disable-sslv2hello,omitempty"`
	HTTPSHost              string         `json:"https-host,omitempty"`
	Sni                    shared.Boolean `json:"sni,omitempty"`
	HTTPSExpect            shared.Boolean `json:"https-expect,omitempty"`
	HTTPSResponseCode      string         `json:"https-response-code,omitempty"`
	HTTPSResponseCodeRegex string         `json:"response-code-regex,omitempty"`
	HTTPSText              string         `json:"https-text,omitempty"`
	TextRegex              string         `json:"text-regex,omitempty"`
	HTTPSURL               string         `json:"https-url,omitempty"`
	URLType                HTTPMethod     `json:"url-type,omitempty"`
	URLPath                string         `json:"url-path,omitempty"`
	PostPath               string         `json:"post-path,omitempty"`
	PostType               string         `json:"post-type,omitempty"`
	HTTPSPostData          string         `json:"https-postdata,omitempty"`
	HTTPSPostfile          string         `json:"https-postfile,omitempty"`
	HTTPSMaintenanceCode   string         `json:"https-maintenance-code,omitempty"`
	Maintenance            shared.Boolean `json:"maintenance,omitempty"`
	MaintenanceText        string         `json:"maintenance-text,omitempty"`
	MaintenanceTextRegex   string         `json:"maintenance-text-regex,omitempty"`
	HTTPSUsername          string         `json:"https-username,omitempty"`
	HTTPSServerCertName    string         `json:"https-server-cert-name,omitempty"`
	HTTPSPassword          string         `json:"https-password,omitempty"`
	HTTPSPasswordString    string         `json:"https-password-string,omitempty"`
	HTTPSKerberosAuth      shared.Boolean `json:"https-kerberos-auth,omitempty"`
	HTTPSKerberosRealm     string         `json:"https-kerberos-realm,omitempty"`
	HTTPSKerberosKdc       KerberosKdc    `json:"https-kerberos-kdc,omitempty"`
	CertKeyShared          shared.Boolean `json:"cert-key-shared,omitempty"`
	Cert                   string         `json:"cert,omitempty"`
	Key                    string         `json:"key,omitempty"`
	KeyPassPhrase          string         `json:"key-pass-phrase,omitempty"`
	KeyPhrase              string         `json:"key-phrase,omitempty"`
}

type TCP struct {
	shared.AxaBase  `json:",inline"`
	MethodTCP       shared.Boolean `json:"method-tcp,omitempty"`
	TCPPort         int            `json:"tcp-port,omitempty"`
	PortHalfopen    shared.Boolean `json:"port-halfopen,omitempty"`
	PortSend        string         `json:"port-send,omitempty"`
	PortResp        string         `json:"port-resp,omitempty"`
	Maintenance     shared.Boolean `json:"maintenance,omitempty"`
	MaintenanceText string         `json:"maintenance-text,omitempty"`
}

type ICMP struct {
	shared.AxaBase `json:",inline"`
	ICMP           shared.Boolean `json:"icmp,omitempty"`
	IP             string         `json:"ip,omitempty"`
	IPV6           string         `json:"ipv6,omitempty"`
	Transparent    shared.Boolean `json:"transparent,omitempty"`
}

type KerberosKdc struct {
	HTTPKerberosHostIP   string `json:"http-kerberos-hostip,omitempty"`
	HTTPKerberosHostIPV6 string `json:"http-kerberos-hostipv6,omitempty"`
	HTTPKerberosPort     int    `json:"http-kerberos-port,omitempty"`
	HTTPKerberosPortV6   int    `json:"http-kerberos-portv6,omitempty"`
}
