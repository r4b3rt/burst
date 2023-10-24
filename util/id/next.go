package id

import "github.com/rs/xid"

func Next() string {
	return xid.New().String()
}
