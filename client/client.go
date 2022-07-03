package client

import (
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/fzdwx/burst/pkg/wsx"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/url"
	"time"
)

type (
	Client struct {
		*wsx.Wsx
		token      string
		config     Config
		serverAddr string
		serverHost string
		proxy      map[string]pkg.ClientProxyInfo
		internet   map[string]*InternetService
	}
)

func NewClient(token string, config Config) *Client {
	return &Client{
		token:      token,
		config:     config,
		serverAddr: burst.FormatAddr(config.Server.Host, config.Server.Port),
		serverHost: config.Server.Host,
		proxy:      make(map[string]pkg.ClientProxyInfo),
		internet:   make(map[string]*InternetService),
	}
}

// Connect  to server,and init connection
func (c *Client) Connect(init func(wsx *wsx.Wsx)) {
	u := url.URL{
		Scheme:   "ws",
		Host:     c.serverAddr,
		Path:     "accept",
		RawQuery: "token=" + c.token,
	}

	// connect to server and check not error
	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		if resp.Body != nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logx.Fatal().Err(err).Msg("connect to server read resp body")
			}
			logx.Fatal().Msgf("connect to server fail: %s", string(body))
		}
		return
	}

	c.Wsx = wsx.NewClassicWsx(conn)

	if init != nil {
		init(c.Wsx)
	}

	go c.Wsx.StartReading(time.Second * 20)
	go c.Wsx.StartWriteHandler(time.Second * 5)
}

// AddProxy add proxy info to client
//
// do not consider repetition
func (c *Client) AddProxy(a protocal.AddProxy) {
	for _, info := range a.Proxy {
		c.proxy[info.Key()] = info
		logx.Info().Msgf("add proxy: intranet [%s] to [%s]", info.IntranetAddr, info.Address(c.serverHost))
	}
}

func (c *Client) GetProxy(key string) (pkg.ClientProxyInfo, bool) {
	info, ok := c.proxy[key]
	return info, ok
}

func (c *Client) AddInterNetService(net *InternetService) {
	c.internet[net.connId] = net
}

func (c Client) GetInternetService(connId string) (*InternetService, bool) {
	internet, ok := c.internet[connId]
	return internet, ok
}