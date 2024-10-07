package service

import (
	"client/model"
	"client/rpc"
	flashkill "client/rpc/kitex_gen/FlashKill"
	"context"
)

func BuyerRegister(b model.Buyers) (aToken string, err error) {
	err = rpc.FlashKill.Register(context.Background(), &flashkill.Buyer{
		Username: b.Name,
		Password: b.Password,
	}, &flashkill.Seller{
		Name: "",
	})
	if err != nil {
		return "", err
	}
	return aToken, nil
}
