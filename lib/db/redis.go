package db

import "github.com/go-redis/redis/v8"

var Kvs *redis.Client

func init() {
	Kvs = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
