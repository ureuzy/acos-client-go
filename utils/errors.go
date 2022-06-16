package utils

import "errors"

func EmptyStringError(s string) error {
	if s == "" {
		return errors.New("must be specified")
	}
	return nil
}
