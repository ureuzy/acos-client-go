package activepartition

type Partition struct {
	CurrentPartitionName string `json:"curr_part_name"`
	Shared               bool   `json:"shared"`
}
