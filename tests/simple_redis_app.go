package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

func main() {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:9093",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// then we can set a Key
	err := rdb.Set(ctx, "index", "0", 0).Err()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1_000_000; i++ {
		result, _ := rdb.Get(ctx, "index").Result()
		v, _ := strconv.Atoi(result)
		v++
		inc := fmt.Sprintf("%d", v)
		log.Println(inc)
		err = rdb.Set(ctx, "index", inc, 0).Err()
	}
}
