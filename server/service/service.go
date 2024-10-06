package service

import (
	"log"
	"server/dao"
	"server/model"
	flashkill "server/rpc/kitex_gen/FlashKill"
	"server/selfUtils"
)

func Register(s *flashkill.Seller, b *flashkill.Buyer) error {
	if b.Username != "" {
		temp, err := selfUtils.Crypto(b.Password)
		b.Password = temp
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err = dao.DB.Model(&model.Buyers{}).Create(&b).Error; err != nil {
			return err
		}
	} else if s.Name != "" {
		temp, err := selfUtils.Crypto(s.Password)
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
