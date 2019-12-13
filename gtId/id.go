package gtId

import "github.com/rs/xid"

func GenId() string {
	return xid.New().String()
}
