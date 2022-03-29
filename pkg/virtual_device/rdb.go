package virtual_device

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var ctx = context.Background()

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PWD"), // no password set
		DB:       0,                      // use default DB
	})
}

func GetLamp() bool {
	val, _ := rdb.Get(ctx, "key").Bool()
	return val
}

func SetLamp(b bool) {
	_ = rdb.Set(ctx, "key", b, 0).Err()
}
