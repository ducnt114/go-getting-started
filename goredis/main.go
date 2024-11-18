package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-getting-started/goredis/conf"
	"time"
)

func main() {
	fmt.Println("Version:", conf.Version)
	fmt.Println("CommitHash:", conf.CommitHash)
	fmt.Println("BuildTimestamp:", conf.BuildTimestamp)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	key := "unique_key"
	for i := 0; i < 2; i++ {
		go func() {
			existed, err := rdb.SetNX(context.Background(),
				key, "random", 30*time.Second).
				Result()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Key:", key, "Existed:", existed)
			}
		}()
	}
	time.Sleep(60 * time.Second)
}
