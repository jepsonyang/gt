package gtRedis

import (
	"errors"
	"fmt"
)

var ErrKeyNotExist = errors.New("key not exist")
var ErrTimeout = errors.New("timeout")

func formatError(err error, format string, v ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, v...)
	return fmt.Errorf("%s originErr: %s", msg, err.Error())
}
