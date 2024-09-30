package model

import (
	"FlashKill/server/dao"
	"time"
)

func (b *Buyers) New() error {
	if err := dao.DB.Model(&Buyers{}).Create(&b).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func (b *Buyers) NewOrder(item Items) Orders {
	return Orders{
		Status:   false,
		BuyerID:  b.BuyerID,
		Item:     item,
		LastTime: time.Now().Add(5 * time.Minute),
	}
}
func (b *Buyers) MyOrder() (Orders, error) {
	var order Orders
	res := dao.DB.Model(&Orders{}).Where("buyer_id=?", b.BuyerID).Find(&order)
	if res.Error != nil {
		return Orders{}, res.Error
	}
	return order, nil
}
