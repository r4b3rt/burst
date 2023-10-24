package serve

import (
	"github.com/fzdwx/burst/api"
	"github.com/gin-gonic/gin"
	"github.com/lxzan/gws"
	"strconv"
	"sync"
)

func ListenAndServe(port int) error {
	s := newServer(port)
	return s.ListenAndServe()
}

type server struct {
	port           int
	connectionLock sync.RWMutex
	connections    map[string]*connection
	wsPort         int64
}

func newServer(port int) *server {
	return &server{
		port:        port,
		connections: map[string]*connection{},
	}
}

func (s *server) ListenAndServe() error {
	upgrader := gws.NewUpgrader(s.e(), &gws.ServerOption{
		ReadAsyncEnabled: true,
		CompressEnabled:  true,
		Recovery:         gws.Recovery,
	})

	engine := gin.New()
	engine.GET("/ws", func(context *gin.Context) {
		socket, err := upgrader.Upgrade(context.Writer, context.Request)
		if err != nil {
			return
		}
		go func() {
			socket.ReadLoop()
		}()
	})

	engine.POST("/export", func(context *gin.Context) {
		var req = &api.ExportRequest{}
		if err := context.ShouldBindJSON(req); err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		resp, err := s.Export(context, req)
		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(200, resp)
	})

	return engine.Run(":" + strconv.Itoa(s.port))
}

func (s *server) e() gws.Event {
	return &q{
		s: s,
	}
}
