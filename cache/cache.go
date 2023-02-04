package cache

import "github.com/bat-labs/krake/pkg/api"

type Cache interface {
	Set(key string, value api.Value)
	Get(key string) api.Value
}

type InMemoryCache struct {
	data map[string]api.Value
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{data: map[string]api.Value{}}
}

func (n *InMemoryCache) Set(key string, value api.Value) {
	n.data[key] = value
}

func (n *InMemoryCache) Get(key string) api.Value {
	return n.data[key]
}
