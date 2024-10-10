package model

import (
	"server/dao"
)

func (b *Buyers) MyOrder() (Orders, error) {
	var order Orders
	res := dao.DB.Model(&Orders{}).Where("buyer_id=?", b.BuyerID).Find(&order)
	if res.Error != nil {
		return Orders{}, res.Error
	}
	return order, nil
}
