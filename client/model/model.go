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
	Item     Items      `gorm:"foreignkey:ItemID"`
}

type Items struct {
	ItemID uint64          `gorm:"primaryKey"`
	Name   string          `json:"name" gorm:"size:255;unique;not null"`
	Price  decimal.Decimal `json:"price" gorm:"type:decimal(10,2)"`
}

type Buyers struct {
	BuyerID  uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Token    string `gorm:"type:varchar(255);not null"`
	Order    Orders `gorm:"foreignkey:OrderID"`
}

type Activities struct {
	ActivityID uint      `gorm:"primaryKey"`
	SellerID   uint      `gorm:"foreignkey:SellerID"`
	Left       uint      `json:"left"`
	LastTime   time.Time `json:"last_time"`
	Item       Items     `gorm:"foreignkey:ItemID"`
}

type Orders struct {
	OrderID  uint      `gorm:"primaryKey"`
	BuyerID  uint      `json:"buyer_id"`
	SellerID uint      `json:"seller_id"`
	Status   bool      `json:"status"`
	Item     Items     `gorm:"foreignkey:ItemID"`
	LastTime time.Time `json:"last_time"`
}
