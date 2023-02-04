package internal

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

const TestPort = 6379

func TestKrake_ListenAndServeCanSetValues(t *testing.T) {
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
}

func TestKrake_ListenAndServeAcksHELLO(t *testing.T) {
	// given a krake server
	k := NewKrakeServer()
	go k.ListenAndServe(fmt.Sprintf("localhost:%d", TestPort))
	defer k.Close()

	// when we set up an actual redis client
	redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
