package cache

import (
	"github.com/fzdwx/burst/pkg"
	"github.com/zeromicro/go-zero/core/collection"
)

var (
	ProxyInfoContainer proxyInfoCache
)

type (
	proxyInfoCache struct {
		m *collection.Cache
	}

	ProxyInfos struct {
		m map[string]*pkg.ServerProxyInfo
	}
)

func NewProxyInfos() *ProxyInfos {
	return &ProxyInfos{m: map[string]*pkg.ServerProxyInfo{}}
}

func (pc *proxyInfoCache) Add(token string) {
	pc.m.Set(token, NewProxyInfos())
}

func (pc proxyInfoCache) Has(token string) bool {
	_, b := pc.m.Get(token)
	return b
}

func (pc *proxyInfoCache) Put(token string, infos []*pkg.ServerProxyInfo) bool {
	old, b := pc.Get(token)
	if !b {
		return false
	}
	old.AddAll(infos)
	return true
}

func (pc proxyInfoCache) Get(token string) (*ProxyInfos, bool) {
	if v, ok := pc.m.Get(token); ok {
		infos := v.(*ProxyInfos)
		return infos, ok
	}

	return nil, false
}

func (pc *proxyInfoCache) Remove(token string) {
	pc.m.Del(token)
}

func (pi *ProxyInfos) Add(info *pkg.ServerProxyInfo) {
	pi.m[info.Addr] = info
}

func (pi *ProxyInfos) AddAll(infos []*pkg.ServerProxyInfo) {
	for _, info := range infos {
		pi.Add(info)
	}
}

func (pi *ProxyInfos) Has(addr string) bool {
	if _, ok := pi.m[addr]; ok {
		return ok
	}
	return false
}