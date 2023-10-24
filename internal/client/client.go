package client

import (
	"bytes"
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
	url := address
	if strings.HasPrefix(address, ":") {
		url = "http://localhost" + address
	}

	c := &client{url}
	return c, nil
}
