package main

import (
	"log"
	"server/rpc"
	"server/rpc/kitex_gen/FlashKill/flashkill"
)

func main() {
	svr := flashkill.NewServer(new(rpc.FlashKillImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
