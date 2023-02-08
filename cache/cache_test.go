package cache

import (
	"github.com/bat-labs/krake/pkg/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMap(t *testing.T) {
	cache := NewInMemoryCache()
	cache.HSet("user_info", "name", api.EncodeBulkString("Felix Angell"))
	cache.HSet("user_info", "age", api.EncodeInteger(24))
	cache.HSet("user_info", "country", api.EncodeBulkString("UK"))

	v, ok := cache.HGet("user_info", "name")
	assert.Equal(t, api.EncodeBulkString("Felix Angell"), v)
	assert.True(t, ok)
}

func TestHashMapNoEntryFound(t *testing.T) {
	cache := NewInMemoryCache()
	v, ok := cache.HGet("pickles", "name")
	assert.Nil(t, v)
	assert.False(t, ok)
}

func TestHashMapDualWrite(t *testing.T) {
	cache := NewInMemoryCache()
	cache.HSet("user_info", "name", api.EncodeBulkString("Felix Angell"))
	cache.HSet("user_info", "name", api.EncodeBulkString("Addy"))

	v, ok := cache.HGet("user_info", "name")
	assert.Equal(t, api.EncodeBulkString("Addy"), v)
	assert.True(t, ok)
}

func TestHashMapFieldDeletion(t *testing.T) {
	cache := NewInMemoryCache()
	cache.HSet("user_info", "name", api.EncodeBulkString("Felix Angell"))
	cache.HSet("user_info", "age", api.EncodeInteger(24))
	cache.HSet("user_info", "country", api.EncodeBulkString("UK"))

	err := cache.HDel("user_info", "name", "country")
	assert.NoError(t, err)

	_, exist := cache.HGet("user_info", "name")
	assert.False(t, exist)

	age, exist := cache.HGet("user_info", "age")
	assert.True(t, exist)
	assert.Equal(t, api.EncodeInteger(24), age)

	_, exist = cache.HGet("user_info", "country")
	assert.False(t, exist)
}

func TestDeletesEntireHashMap(t *testing.T) {
	cache := NewInMemoryCache()
	cache.HSet("user_info", "name", api.EncodeBulkString("Felix Angell"))
	cache.HSet("user_info", "age", api.EncodeInteger(24))
	cache.HSet("user_info", "country", api.EncodeBulkString("UK"))

	err := cache.HDel("user_info")
	assert.NoError(t, err)

	_, exist := cache.HGet("user_info", "name")
	assert.False(t, exist)

	_, exist = cache.HGet("user_info", "age")
	assert.False(t, exist)

	_, exist = cache.HGet("user_info", "country")
	assert.False(t, exist)
}
