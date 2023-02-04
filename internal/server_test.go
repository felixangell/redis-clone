package internal

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

const TestPort = 6379

func TestKrake_ListenAndServeCanSetAndRetriveValues(t *testing.T) {
	// given a krake server
	k := NewKrakeServer()
	go k.ListenAndServe(fmt.Sprintf("localhost:%d", TestPort))
	defer k.Close()

	// when we set up an actual redis client
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// then we can set a Key
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// and retrieve it.
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	rdb.Close()
}

func TestKrake_ListenAndServeAcksHELLO(t *testing.T) {
	// given a krake server
	k := NewKrakeServer()
	go k.ListenAndServe(fmt.Sprintf("localhost:%d", TestPort))
	defer k.Close()

	// when we set up an actual redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Close()
	assert.NoError(t, err)
}
