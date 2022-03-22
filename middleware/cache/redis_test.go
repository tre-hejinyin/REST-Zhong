package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// v8
func ExampleInitClient() {
	ctx := context.Background()
	if err := InitClient(); err != nil {
		return
	}

	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
