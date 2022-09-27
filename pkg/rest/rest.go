package rest

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/axapi/errors"
	"github.com/ureuzy/acos-client-go/utils"
)

// Operator Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/operations.html
type Operator[object any, objectList any] interface {
	Get(names ...string) (*object, error)
	List(names ...string) (*objectList, error)
	Create(object *object, names ...string) (*object, error)
	CreateList(object *objectList, names ...string) (*objectList, error)
	Modify(object *object, names ...string) (*object, error)
	Replace(object *object, names ...string) (*object, error)
	Delete(names ...string) error
}

type operator[object any, objectList any] struct {
	utils.HTTPClient
	basePath string
}

func Rest[object any, objectList any](c utils.HTTPClient, path string) Operator[object, objectList] {
	return &operator[object, objectList]{HTTPClient: c, basePath: path}
}

func (o *operator[object, objectList]) Get(names ...string) (*object, error) {
	err := errors.EmptyStringArrayError(names)
	if err != nil {
		return nil, err
	}

	res, err := o.GET(o.getFormattedURL(o.basePath+"/%s", names))
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response object
	if err = res.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (o *operator[object, objectList]) List(names ...string) (*objectList, error) {
	res, err := o.GET(o.getFormattedURL(o.basePath, names))
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response objectList
	if err = res.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (o *operator[object, objectList]) Create(instance *object, names ...string) (*object, error) {
	res, err := o.POST(o.getFormattedURL(o.basePath, names), instance)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response object
	if err = res.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (o *operator[object, objectList]) CreateList(instance *objectList, names ...string) (*objectList, error) {
	res, err := o.POST(o.getFormattedURL(o.basePath, names), instance)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response objectList
	if err = res.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (o *operator[object, objectList]) Modify(instance *object, names ...string) (*object, error) {
	err := errors.EmptyStringArrayError(names)
	if err != nil {
		return nil, err
	}

	res, err := o.POST(o.getFormattedURL(o.basePath+"/%s", names), instance)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response object
	if err = res.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (o *operator[object, objectList]) Replace(instance *object, names ...string) (*object, error) {
	err := errors.EmptyStringArrayError(names)
	if err != nil {
		return nil, err
	}

	res, err := o.PUT(o.getFormattedURL(o.basePath+"/%s", names), instance)
	if err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, errors.Handle(res)
	}

	var response object
	if err = res.UnmarshalJSON(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (o *operator[object, objectList]) Delete(names ...string) error {
	err := errors.EmptyStringArrayError(names)
	if err != nil {
		return err
	}

	res, err := o.DELETE(o.getFormattedURL(o.basePath+"/%s", names))
	if err != nil {
		return err
	}

	if res.HasError() {
		return errors.Handle(res)
	}

	return nil
}

func (o *operator[object, objectList]) getFormattedURL(urlFormat string, args []string) string {
	is := make([]interface{}, len(args))
	for i := range args {
		is[i] = args[i]
	}
	return fmt.Sprintf(urlFormat, is...)
}
