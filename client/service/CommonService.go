package service

import (
	"client/model"
	"client/rpc"
	flashkill "client/rpc/kitex_gen/FlashKill"
	"context"
	"fmt"
	"log"
)

func BuyerRegister(b model.Buyers) (aToken string, err error) {
	err = rpc.FlashKill.Register(context.Background(), &flashkill.Buyer{
		Username: b.Name,
		Password: b.Password,
	}, &flashkill.Seller{
		Name: "",
	})
	if err != nil {
		log.Fatalln("register rpc Buyer error:", err)
	}
	aToken, err = rpc.FlashKill.GenToken(context.Background(), &flashkill.Buyer{
		Username: b.Name,
		Password: b.Password,
	}, &flashkill.Seller{
		Name: "",
	})
	if err != nil {
		log.Fatalln("Buyer set token error:", err)
	}
	return
}
func SellerRegister(s model.Sellers) (aToken string, err error) {
	err = rpc.FlashKill.Register(context.Background(), &flashkill.Buyer{
		Username: "",
	}, &flashkill.Seller{
		Name:     s.Name,
		Password: s.Password,
	})
	if err != nil {
		log.Fatalln("register rpc Seller error:", err)
	}
	aToken, err = rpc.FlashKill.GenToken(context.Background(), &flashkill.Buyer{
		Username: "",
	}, &flashkill.Seller{
		Name:     s.Name,
		Password: s.Password,
	})
	if err != nil {
		log.Fatalln("Seller set token error:", err)
	}
	return
}
func SellerLogin(s model.Sellers) (model.Sellers, error) {
	err := rpc.FlashKill.Login(context.Background(), &flashkill.Buyer{
		Username: "",
	}, &flashkill.Seller{
		Name:     s.Name,
		Password: s.Password,
	})
	if err != nil {
		fmt.Println("failed")
		return model.Sellers{}, err
	}
	return model.Sellers{}, nil
}
func BuyerLogin(b model.Buyers) (err error) {
	err = rpc.FlashKill.Login(context.Background(), &flashkill.Buyer{
		Token: b.Token,
	}, &flashkill.Seller{
		Name: "",
	})
	if err != nil {
		log.Fatalln("Buyer login error :", err)
	}
	return nil
}
