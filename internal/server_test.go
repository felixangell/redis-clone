package internal

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net"
	"testing"
	"time"
)

func generateTestPort() int {
	rand.Seed(time.Now().UnixNano())

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	return ln.Addr().(*net.TCPAddr).Port
}

func TestKrake_ListenAndServeCanSetAndRetrieveValues(t *testing.T) {
	testPort := generateTestPort()

	// given a krake server
	k := NewKrakeServer()
	go k.ListenAndServe(fmt.Sprintf("localhost:%d", testPort))
	defer k.Close()

	// when we set up an actual redis client
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", testPort),
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
	testPort := generateTestPort()

	// given a krake server
	k := NewKrakeServer()
	go k.ListenAndServe(fmt.Sprintf("localhost:%d", testPort))
	defer k.Close()

	// when we set up an actual redis client
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", testPort),
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
	testPort := generateTestPort()

	// given a krake server
	k := NewKrakeServer()
	go k.ListenAndServe(fmt.Sprintf("localhost:%d", testPort))
	defer k.Close()

	// when we set up an actual redis client
	redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", testPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
