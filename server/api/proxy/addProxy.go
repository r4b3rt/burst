package proxy

import (
	"errors"
	"fmt"
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/fzdwx/burst/pkg/result"
	"github.com/fzdwx/burst/server/cache"
	"github.com/fzdwx/burst/server/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net"
	"net/http"
)

func AddProxy(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err, req, token := addProxyPreCheck(r)
		if err != nil {
			result.HttpBadRequest(w, err.Error())
			return
		}

		info, ok := cache.ProxyInfoContainer.Get(token)
		if !ok {
			result.HttpBadRequest(w, proxyInfoNotFound.Error())
		}

		server, b := cache.ServerContainer.Get(token)
		if !b {
			result.HttpBadRequest(w, clientNotFound.Error())
			return
		}

		var proxyInfos []*pkg.ServerProxyInfo

		// check if proxy is duplicated
		for _, proxyInfo := range req.Proxy {
			if info.Has(proxyInfo.addr()) {
				result.HttpBadRequest(w, fmt.Sprintf("proxy %s already exists", proxyInfo.String()))
				return
			}
			proxyInfos = append(proxyInfos, proxyInfo.toCache())
		}

		err, clientProxyInfos, closers := server.Lunch(proxyInfos)
		clean := func() {
			for _, c := range closers {
				c.Close()
			}
		}

		if err != nil {
			go clean()
			result.HttpBadRequest(w, err.Error())
			return
		}

		bytes, err := protocal.NewAddProxy(clientProxyInfos).Encode()
		if err != nil {
			go clean()
			result.HttpBadRequest(w, err.Error())
			return
		}

		// send client proxy info to client
		server.WriteBinary(bytes)

		cache.ProxyInfoContainer.Put(token, proxyInfos)

		httpx.OkJson(w, clientProxyInfos)
	}
}

type (
	addProxyInfoReq struct {
		Proxy []addProxyInfo `json:"proxy"`
	}

	addProxyInfo struct {
		Ip          string `json:"ip"`
		Port        int    `json:"port"`
		ChannelType string `json:"channelType,default=tcp"`
	}
)

var (
	proxyIsRequired   = errors.New("proxy is required")
	tokenIsRequired   = errors.New("token is required")
	proxyInfoNotFound = errors.New("the proxy info not found")
	ipIsBlank         = errors.New("ip is blank")
	ipIsNotValid      = errors.New("ip is not valid")
	portIsNotValid    = errors.New("port is not valid")
	clientNotFound    = errors.New("client not found")
)

func (i addProxyInfo) addr() string {
	return burst.FormatAddr(i.Ip, i.Port)
}

func (i addProxyInfo) String() string {
	return fmt.Sprintf("%s:%s", i.ChannelType, i.addr())
}

func (i addProxyInfo) toCache() *pkg.ServerProxyInfo {
	return &pkg.ServerProxyInfo{
		Ip:          i.Ip,
		Port:        i.Port,
		ChannelType: i.ChannelType,
		Addr:        i.addr(),
	}
}

func addProxyPreCheck(r *http.Request) (error, *addProxyInfoReq, string) {
	token := burst.GetPars("token", r)
	if token == burst.EmptyStr {
		return tokenIsRequired, nil, ""
	}

	var req addProxyInfoReq
	err := httpx.ParseJsonBody(r, &req)
	if err != nil {
		return err, nil, ""
	}

	err = req.check()
	if err != nil {
		return err, nil, ""
	}
	return nil, &req, token
}

func (r addProxyInfoReq) check() error {
	if len(r.Proxy) == 0 {
		return proxyIsRequired
	}

	for _, info := range r.Proxy {
		if info.Ip == burst.EmptyStr {
			return ipIsBlank
		}

		if info.Ip != "localhost" {
			ip := net.ParseIP(info.Ip)
			if ip == nil {
				return ipIsNotValid
			}
		}

		if info.Port < 0 || info.Port > 65535 {
			return portIsNotValid
		}

		if !pkg.CheckChannelType(info.ChannelType) {
			return pkg.ErrChannelTypeNotValid
		}
	}
	return nil
}