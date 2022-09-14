package activepartition

import "github.com/ureuzy/acos-client-go/pkg/axapi/shared"

type Partition struct {
	CurrentPartitionName string         `json:"curr_part_name"`
	Shared               shared.Boolean `json:"shared"`
}
