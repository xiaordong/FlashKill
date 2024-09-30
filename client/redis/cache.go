package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var RDB *redis.Client

func InitRedis() {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	if err := db.Ping().Err(); err != nil {
		log.Fatal(err)
	}
	RDB = db
}
