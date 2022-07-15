package proxy

import (
	"github.com/fzdwx/burst/pkg"
	"io"
)

var AloneHttpServer = aloneHttpServer{}

type (
	AloneHttpServerConfig struct {
		Port   int  `json:",default=39939"`
		Enable bool `json:",default=true"`
	}

	aloneHttpServer struct {
	}
)

func (ahs aloneHttpServer) Run(config AloneHttpServerConfig) {
	// todo
}

// handlerHttp.
func (c *Container) handlerHttp(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	// todo handler http
	return nil, nil, nil
}
