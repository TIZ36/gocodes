package service

import (
	"encoding/json"
	"fmt"
	"gomock/common"
	"reflect"
)

type CacheService interface {
	CacheIn(string, string, any)
	CacheOut(string, string) any
	Register(string, reflect.Type)
}

type NormalService struct {
	cache             common.CommonCache[string, []byte]
	funcReturnTypeMap map[string]reflect.Type
}

func key(label string, k string) string {
	return fmt.Sprintf("%s-%s", label, k)
}

func (queryService *NormalService) CacheIn(api string, k string, v any) {
	valueBin, _ := json.Marshal(v)
	cacheKey := key(api, k)
	_ = queryService.cache.Put(cacheKey, valueBin)
}

func (queryService *NormalService) CacheOut(api string, k string) any {
	v, err := queryService.cache.Get(key(api, k))

	if err != nil {
		return nil
	}

	return v
}

func (queryService *NormalService) Register(api string, returnType reflect.Type) {
	queryService.funcReturnTypeMap[api] = returnType
}
