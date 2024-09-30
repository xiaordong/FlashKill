package main

import (
	"FlashKill/server/dao"
	"FlashKill/server/model"
	"log"
)

func main() {
	dao.InitMySQL()
	err := dao.DB.AutoMigrate(&model.Sellers{}, &model.Buyers{}, &model.Orders{}, &model.Items{}, &model.Activities{})
	if err != nil {
		log.Fatal(err)
	}
}
