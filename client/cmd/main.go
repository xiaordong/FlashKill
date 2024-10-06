package main

import (
	"client/redis"
	"client/router"
	"client/rpc"
)

func main() {
	redis.InitRedis()
	rpc.InitRpc()
	router.InitRouter()
}
