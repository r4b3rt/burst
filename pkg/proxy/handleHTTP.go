package proxy

import (
	"github.com/fzdwx/burst/pkg"
	"io"
)

var AloneHttpServer = aloneHttpServer{
	enable: false,
}

type (
	AloneHttpServerConfig struct {
		Port   int  `json:",default=39939"`
		Enable bool `json:",default=true"`
	}

	aloneHttpServer struct {
		enable bool
	}
)

func (ahs *aloneHttpServer) Run(config AloneHttpServerConfig) {
	ahs.enable = config.Enable
	if !ahs.enable {
		return
	}

	// todo startup alone http server
}

// handlerHttp.
func (c *Container) handlerHttp(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	// todo
	//  1. save info to cache
	return nil, nil, nil
}
