package gtRedis

import (
	"errors"
	"fmt"
)

func formatError(err error, format string, v ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, v...)
	return errors.New(fmt.Sprintf("%s originErr: %s", msg, err.Error()))
}
