package datasource

import (
    "log"
    "github.com/go-redis/redis"
)

var Redis *redis.Client

const (
    addr = "localhost:6379"
    pass = ""
    db   = 0
)

func SetupRedis() {
    Redis = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: pass,
        DB:       db,
    })
    _, err := Redis.Ping().Result()

    if err != nil {
        log.Fatal(err)
    }
}