package redis

import (
	"client/model"
	"github.com/go-redis/redis"
	"log"
)

func InitRedis() {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	if err := db.Ping().Err(); err != nil {
		log.Fatal(err)
	}
	model.RDB = db
}
