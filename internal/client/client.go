package client

import (
	"bytes"
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/jsonutil"
	"io"
	"net/http"
	"strings"
)

type client struct {
	url string
}

func (c client) Export(a *api.ExportRequest) (*api.ExportResponse, error) {
	resp, err := http.Post(c.url, "application/json", strings.NewReader(string(jsonutil.Encode(a))))
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return jsonutil.Decode2(bytes.NewBuffer(b)), nil
}

func New(address string) (*client, error) {
	c := &client{fmt.Sprintf("http://%s/export", address)}
	return c, nil
}
