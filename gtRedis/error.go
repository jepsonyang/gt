package gtRedis

import (
	"errors"
	"fmt"
)

var ErrKeyNotExist = errors.New("key not exist")
var ErrTimeout = errors.New("timeout")
var ErrLockExist = errors.New("lock already exist")
var ErrLockNotExist = errors.New("lock not exist")

func formatError(err error, format string, v ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, v...)
	return fmt.Errorf("%s originErr: %s", msg, err.Error())
}
