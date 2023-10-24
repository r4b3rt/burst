package jsonutil

import (
	"bytes"
	"encoding/json"
	"github.com/fzdwx/burst/api"
)

func Encode(v any) []byte {
	marshal, _ := json.Marshal(v)
	return marshal
}

func Decode(buf *bytes.Buffer) *api.TransFromData {
	var v = &api.TransFromData{}
	_ = json.NewDecoder(buf).Decode(v)
	return v
}
