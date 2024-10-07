package model

import (
	"log"
	"server/dao"
	"time"
)

// NewActivity 商家创建活动,需要选择商品和持续时间
func (s *Sellers) NewActivity(item Items, lastTime time.Time) (Activities, error) {
	a := Activities{
		SellerID: s.SellerID,
		Item:     item,
		LastTime: lastTime,
	}
	res := dao.DB.Model(&Activities{}).Create(&a)
	if res.Error != nil {
		log.Println(res.Error)
		return a, res.Error
	}
	return a, nil
}

// GetOrders 商家获取订单
func (s *Sellers) GetOrders() ([]Orders, error) {
	var orders []Orders
	if err := dao.DB.Preload("Buyers").Preload("Items").Where("seller_id = ?", s.SellerID).Find(&orders).Error; err != nil {
		return orders, err
	}
	return orders, nil
}
