package main

import (
	"FlashKill/client/redis"
	"FlashKill/client/router"
)

func main() {
	redis.InitRedis()
	router.InitRouter()
}
