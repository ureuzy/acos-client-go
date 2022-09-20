package rest

import (
	"fmt"

	"github.com/ureuzy/acos-client-go/pkg/axapi/errors"
	"github.com/ureuzy/acos-client-go/utils"
)

// Operator Docs: https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/operations.html
type Operator[object any, objectList any] interface {
	Get(name string, parentnames ...string) (*object, error)
	List(parentnames ...string) (*objectList, error)
	Create(object *object, parentnames ...string) (*object, error)
	CreateList(object *objectList, parentnames ...string) (*objectList, error)
	Modify(name string, object *object, parentnames ...string) (*object, error)
	Delete(name string, parentnames ...string) error
}

type operator[object any, objectList any] struct {
	utils.HTTPClient
	basePath string
}

func Rest[object any, objectList any](c utils.HTTPClient, path string) Operator[object, objectList] {
	return &operator[object, objectList]{HTTPClient: c, basePath: path}
}

func (o *operator[object, objectList]) Get(name string, parentnames ...string) (*object, error) {
	err := errors.EmptyStringError(name)
	if err != nil {
		return nil, err
	}

	parentnames = append(parentnames, name)

	res, err := o.GET(o.getFormattedURL(o.basePath+"/%s", parentnames))
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

func (o *operator[object, objectList]) List(parentnames ...string) (*objectList, error) {
	res, err := o.GET(o.getFormattedURL(o.basePath, parentnames))
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

func (o *operator[object, objectList]) Create(instance *object, parentnames ...string) (*object, error) {
	res, err := o.POST(o.getFormattedURL(o.basePath, parentnames), instance)
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

func (o *operator[object, objectList]) CreateList(instance *objectList, parentnames ...string) (*objectList, error) {
	res, err := o.POST(o.getFormattedURL(o.basePath, parentnames), instance)
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

func (o *operator[object, objectList]) Modify(name string, instance *object, parentnames ...string) (*object, error) {
	err := errors.EmptyStringError(name)
	if err != nil {
		return nil, err
	}

	parentnames = append(parentnames, name)

	res, err := o.POST(o.getFormattedURL(o.basePath+"/%s", parentnames), instance)
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

func (o *operator[object, objectList]) Delete(name string, parentnames ...string) error {
	err := errors.EmptyStringError(name)
	if err != nil {
		return err
	}

	parentnames = append(parentnames, name)

	res, err := o.DELETE(o.getFormattedURL(o.basePath+"/%s", parentnames))
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
