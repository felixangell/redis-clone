package cache

import (
	"github.com/bat-labs/krake/pkg/api"
)

type Cache interface {
	Set(key string, value api.Value)
	Get(key string) api.Value
	Del(key string) error
	IncrBy(key string, value api.Value) api.Value
	Exists(key string) bool

	Expire(key string, seconds string, kind api.Value) api.Value

	HSet(hash string, field string, value api.Value)
	HGet(hash string, field string) (api.Value, bool)
	HDel(hash string, fields ...string) error
}

type KMap map[string]api.Value

type InMemoryCache struct {
	keys  KMap
	dicts map[string]*KMap
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		keys:  map[string]api.Value{},
		dicts: map[string]*KMap{},
	}
}

func (n *InMemoryCache) IncrBy(key string, other api.Value) api.Value {
	v, ok := n.keys[key]
	if !ok {
		result := api.EncodeInteger(0)
		n.Set(key, result)
		return result
	}

	increased := v.IncrBy(other)
	n.Set(key, increased)
	return increased
}

func (n *InMemoryCache) Exists(field string) bool {
	_, ok := n.keys[field]
	return ok
}

func (n *InMemoryCache) Set(key string, value api.Value) {
	n.keys[key] = value
}

func (n *InMemoryCache) Get(key string) api.Value {
	return n.keys[key]
}

func (n *InMemoryCache) Del(key string) error {
	delete(n.keys, key)
	return nil
}

func (n *InMemoryCache) HSet(hash string, field string, value api.Value) {
	_, exist := n.dicts[hash]
	if !exist {
		n.dicts[hash] = &KMap{}
	}
	(*n.dicts[hash])[field] = value
}

func (n *InMemoryCache) HGet(hash string, field string) (api.Value, bool) {
	d, exist := n.dicts[hash]
	if !exist {
		return nil, false
	}
	v, ok := (*d)[field]
	return v, ok
}

func (n *InMemoryCache) HDel(hash string, fields ...string) error {
	if len(fields) == 0 {
		// delete the whole map
		delete(n.dicts, hash)
		return nil
	}

	d, ok := n.dicts[hash]
	if !ok {
		return nil
	}

	for _, f := range fields {
		delete(*d, f)
	}
	return nil
}

func (n *InMemoryCache) Expire(key string, seconds string, kind api.Value) api.Value {
	//TODO implement me
	panic("implement me")
}
