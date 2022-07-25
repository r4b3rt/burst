package proxy

import (
	"fmt"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/logx"
	"io"
	"net"
)

var AloneHttpServer = aloneHttpServer{
	enable:  false,
	running: false,
}

type (
	AloneHttpServerConfig struct {
		Port   int  `json:",default=39939"`
		Enable bool `json:",default=true"`
	}

	aloneHttpServer struct {
		enable  bool
		running bool
	}
)

// Startup alone http server
func (ahs *aloneHttpServer) Startup(config AloneHttpServerConfig) {
	// ensure aloneHttpServer not running.
	if ahs.running {
		panic("repeated calls aloneHttpServer#Startup")
	}

	// update status
	ahs.running = true
	ahs.enable = config.Enable
	if !ahs.enable {
		return
	}

	// bind port
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		panic(fmt.Errorf("aloneHttpServer resolve tcp adder fail:%v", err))
	}

	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(fmt.Errorf("aloneHttpServer bind port fail:%v", err))
	}

	// accept user request.
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			logx.Err(err).Msg("aloneHttpServer accept error")
		}
		go func() {
			for {
				bytes := make([]byte, 1024)
				n, err := conn.Read(bytes)
				if err != nil {
					logx.Err(err).Msg("aloneHttpServer  error")
				}

				if n == 0 {
					continue
				}
				logx.Info().Msgf("from client get message:%s", string(bytes[:n]))
			}
		}()
	}
}

// handlerHttp.
func (c *Container) handlerHttp(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	// todo
	//  1. save info to cache
	return nil, nil, nil
}
