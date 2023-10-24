package id

import "github.com/rs/xid"

var guid = xid.New()

func Next() string {
	return guid.String()
}
