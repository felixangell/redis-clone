package internal

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

const TestPort = 6379

func TestKrake_ListenAndServeCanSetAndRetrieveValuesInMultipleWrites(t *testing.T) {
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
	err := rdb.Set(ctx, "index", "0", 0).Err()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		result := rdb.Get(ctx, "aadilah_age")
		val, err := result.Result()
	}
}

func TestKrake_ListenAndServeCanSetAndRetrieveValues(t *testing.T) {
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
	err := rdb.Set(ctx, "aadilah_age", "23", 0).Err()
	if err != nil {
		panic(err)
	}

	result := rdb.Get(ctx, "aadilah_age")
	val, err := result.Result()
	assert.NoError(t, err)
	assert.Equal(t, "23", val)
}

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
	err := rdb.Set(ctx, "aadilah_age", "23", 0).Err()
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
