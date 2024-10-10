package model

import (
	"server/dao"
)

// GetOrders 商家获取订单
func (s *Sellers) GetOrders() ([]Orders, error) {
	var orders []Orders
	if err := dao.DB.Preload("Buyers").Preload("Items").Where("seller_id = ?", s.SellerID).Find(&orders).Error; err != nil {
		return orders, err
	}
	return orders, nil
}
