package common

import (
	"encoding/json"
	"fmt"
)

type DbCache[K any, V *any] interface {
	Get(K) (V, error)
	Put(K, V) error
	Del(K) error
	Update(K, V) error
}

// CommonCache 通用缓存结构 /
type CommonCache[K any, V any] struct {
}

func (commonCache *CommonCache[K, V]) Get(k any) (any, error) {

	var re V
	key := Key(k)
	cacheRe, err := AppCtx.InMemoryCache.Get(key)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(cacheRe, &re)

	fmt.Println("from cache:", re)
	return re, nil
}

func (commonCache *CommonCache[K, V]) Put(key K, value V) error {
	kStr := Key(key)
	vBytes, _ := json.Marshal(value)

	AppCtx.InMemoryCache.Set(kStr, vBytes)
	return nil
}

func (commonCache *CommonCache[K, V]) Del(key K) error {
	return nil
}
func (commonCache *CommonCache[K, V]) Update(key K, value V) error {
	return nil
}

func Key(k any) string {
	kBytes, _ := json.Marshal(k)
	return string(kBytes)
}
