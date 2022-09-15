package shared

import (
	"encoding/json"
)

type AxaBase struct {
	UUID    string `json:"uuid,omitempty"`
	UserTag string `json:"user-tag,omitempty"`
}

type SamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

type Boolean bool

func (b *Boolean) UnmarshalJSON(data []byte) error {
	var i int
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}
	*b = i > 0
	return nil
}
