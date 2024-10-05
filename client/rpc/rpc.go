package rpc

import (
	flashkill "client/rpc/kitex_gen/FlashKill/flashkill"
	"github.com/cloudwego/kitex/client"
	"log"
)

var FlashKill flashkill.Client

func InitRpc() {
	f, err := flashkill.NewClient("uh", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Fatalln("can't connent service :", err)
	}
	FlashKill = f
}
