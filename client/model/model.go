package model

import (
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
)

var RDB *redis.Client

type Sellers struct {
	SellerID uint       `json:"id" gorm:"primaryKey"`
	Name     string     `json:"name" gorm:"size:255;unique;not null"`
	Password string     `json:"password" gorm:"size:255;not null"`
	Token    string     `json:"token" gorm:"size:255;not null"`
	Activity Activities `gorm:"foreignkey:ActivityID"`
}

type Buyers struct {
	BuyerID  uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Token    string `gorm:"type:varchar(255);not null"`
	Order    Orders `gorm:"foreignkey:OrderID"`
}

type Activities struct {
	GoodsName string          `json:"goods_name"`
	Price     decimal.Decimal `json:"price" `
	Left      uint            `json:"left"`
	TimeOut   int             `json:"time_out"`
}

type Orders struct {
	OrderID  uint `gorm:"primaryKey"`
	BuyerID  uint `json:"buyer_id"`
	SellerID uint `json:"seller_id"`
	Status   bool `json:"status"`
	TimeOut  int  `json:"time_out"`
}
