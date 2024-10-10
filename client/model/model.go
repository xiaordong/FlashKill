package model

import (
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
	"time"
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
	ActivityID uint            `gorm:"primaryKey"`
	SellerID   uint            `gorm:"foreignkey:SellerID"`
	Price      decimal.Decimal `json:"price" gorm:"type:decimal(10,2)"`
	Left       uint            `json:"left"`
	LastTime   time.Time       `json:"last_time"`
}

type Orders struct {
	OrderID  uint      `gorm:"primaryKey"`
	BuyerID  uint      `json:"buyer_id"`
	SellerID uint      `json:"seller_id"`
	Status   bool      `json:"status"`
	LastTime time.Time `json:"last_time"`
}
