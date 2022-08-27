package shared

type AxaBase struct {
	UUID    string `json:"uuid,omitempty"`
	UserTag string `json:"user-tag,omitempty"`
}

type SamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}
