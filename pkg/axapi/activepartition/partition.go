package activepartition

import (
	"github.com/ureuzy/acos-client-go/pkg/axapi/errors"
	"github.com/ureuzy/acos-client-go/utils"
)

// Docs: https://acos.docs.a10networks.com/axapi/521p2/axapiv3/active_partition.html
// URI: /axapi/v3/active-partition

type operator struct {
	utils.HTTPClient
	basePath string
}

type Operator interface {
	Set(object Partition) error
}

func New(c utils.HTTPClient) Operator {
	const path = "active-partition"
	return &operator{HTTPClient: c, basePath: path}
}

func (o *operator) Set(object Partition) error {
	err := errors.EmptyStringError(object.CurrentPartitionName)
	if err != nil {
		return err
	}

	res, err := o.POST(o.basePath, object)
	if err != nil {
		return err
	}

	if res.HasError() {
		return errors.Handle(res)
	}

	return nil
}
