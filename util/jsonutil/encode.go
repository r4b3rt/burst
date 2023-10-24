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

func Decode2(buf *bytes.Buffer) *api.ExportResponse {
	var v = &api.ExportResponse{}
	_ = json.NewDecoder(buf).Decode(v)
	return v
}
