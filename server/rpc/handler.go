package rpc

import (
	"context"
	flashkill "server/rpc/kitex_gen/FlashKill"
	"server/service"
)

// FlashKillImpl implements the last service interface defined in the IDL.
type FlashKillImpl struct{}

// Register implements the FlashKillImpl interface.
func (s *FlashKillImpl) Register(ctx context.Context, b *flashkill.Buyer, seller *flashkill.Seller) (err error) {

	return service.Register(seller, b)
}

// Login implements the FlashKillImpl interface.
func (s *FlashKillImpl) Login(ctx context.Context, b *flashkill.Buyer, seller *flashkill.Seller) (err error) {

	return service.Login(seller, b)
}

// GenToken implements the FlashKillImpl interface.
func (s *FlashKillImpl) GenToken(ctx context.Context, b *flashkill.Buyer, seller *flashkill.Seller) (resp string, err error) {
	resp, err = service.SetToken(seller, b)
	if err != nil {
		return
	}
	err = service.WriteToken(b, seller, resp)
	return
}
