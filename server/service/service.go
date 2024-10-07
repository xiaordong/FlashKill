package service

import (
	"errors"
	"log"
	"server/dao"
	"server/model"
	flashkill "server/rpc/kitex_gen/FlashKill"
	"server/utils"
)

func Register(s *flashkill.Seller, b *flashkill.Buyer) error {
	if b.Name != "" {
		temp, err := utils.Crypto(b.Password)
		b.Password = temp
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err = dao.DB.Create(&b).Error; err != nil {
			return err
		}
	} else if s.Name != "" {
		temp, err := utils.Crypto(s.Password)
		s.Password = temp
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err = dao.DB.Model(&model.Sellers{}).Create(&s).Error; err != nil {
			return err
		}
	}
	return nil
}
func Login(s model.Sellers, b model.Buyers) {
	if b.Name != "" {
	} else if s.Name != "" {
	}
}
func GenToken(s *flashkill.Seller, b *flashkill.Buyer) (token string, err error) {
	if b.Name != "" {
		token, err = utils.GenToken(*b.BuyerID)
		if err != nil {
			return "", err
		}
		return token, nil
	} else if s.Name != "" {
		token, err = utils.GenToken(*s.SellerID)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", errors.New("empty data")
}
