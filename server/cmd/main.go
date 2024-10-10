package main

import (
	"log"
	"server/dao"
	"server/model"
	"server/rpc"
	"server/rpc/kitex_gen/FlashKill/flashkill"
)

func main() {
	dao.DB.AutoMigrate(&model.Sellers{}, &model.Buyers{}, &model.Activities{}, &model.Orders{})
	svr := flashkill.NewServer(new(rpc.FlashKillImpl))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}

}
