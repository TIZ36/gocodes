package service

import (
	"encoding/json"
	"fmt"
	"gomock/common"
)

type Service struct {
	funcMap map[string]func()
}

type ProxyService struct {
	Cache  Cache
	Method func() (any, error)
}

type Cache struct {
	Enable   bool
	Cache    common.CommonCache[string, []byte]
	QueryKey string
	label    string
	k        string
	others   string
}

type ServiceContext struct {
}

func (service *Service) Proxy(f func() (any, error)) *ProxyService {
	return &ProxyService{
		Method: f,
	}
}

func (proxyService *ProxyService) Cached(keyCompos ...string) *ProxyService {
	apiKey := buildKey(proxyService.Cache.label, proxyService.Cache.k)

	keyCompos = append(keyCompos, apiKey)

	qk := buildKey(keyCompos...)

	proxyService.Cache.Enable = true
	proxyService.Cache.QueryKey = qk
	proxyService.Cache.Cache = common.CommonCache[string, []byte]{}
	return proxyService
}

func (proxyService *ProxyService) Exec() (any, error) {
	var reBin []byte
	if proxyService.Cache.Enable {
		reBin, _ := proxyService.Cache.Cache.Get(proxyService.Cache.QueryKey)

		if reBin != nil {
			return reBin, nil
		}
	}

	re, err := proxyService.Method()

	if proxyService.Cache.Enable {
		reBin, _ = json.Marshal(re)
		proxyService.Cache.Cache.Put(proxyService.Cache.QueryKey, reBin)
	}

	if err != nil {
		return nil, err
	}

	return reBin, nil
}

func buildKey(compo ...string) string {
	k := ""

	for _, c := range compo {
		k += fmt.Sprintf("%s-", c)
	}

	return k[:len(k)-1]
}
