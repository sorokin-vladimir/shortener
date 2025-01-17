package utils

import (
	"errors"
)

func CheckHttpMethod(rMethod string, allowedMethod string) error {
	if rMethod != allowedMethod {
		errText := "Only " + allowedMethod + " method is allowed"
		return errors.New(errText)
	}

	return nil
}
