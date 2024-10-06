package model

import (
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
	"time"
)

var RDB *redis.Client

type Sellers struct {
	ID       uint       `json:"id" gorm:"primarykey"`
	Name     string     `json:"name" gorm:"size:255;unique;not null"`
	Password string     `json:"password" gorm:"size:255;not null"`
	Activity Activities `gorm:"foreignkey:ActivityID"`
	Item     Items      `gorm:"foreignkey:ItemID"`
}

type Items struct {
	ItemID uint64          `gorm:"primarykey"`
	Price  decimal.Decimal `json:"price" gorm:"type:decimal(10,2)"`
}

type Buyers struct {
	BuyerID  uint   `gorm:"primarykey"`
	Name     string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Order    Orders `gorm:"foreignkey:OrderID"`
}

type Activities struct {
	ActivityID uint      `gorm:"primarykey"`
	PosterID   uint      `json:"poster_id"`
	Left       uint      `json:"left"`
	LastTime   time.Time `json:"last_time"`
	Item       Items     `gorm:"foreignkey:ItemID"`
}

type Orders struct {
	OrderID  uint      `gorm:"primarykey"`
	BuyerID  uint      `json:"buyer_id"`
	SellerID uint      `json:"seller_id"`
	Status   bool      `json:"status"`
	Item     Items     `gorm:"foreignkey:ItemID"`
	LastTime time.Time `json:"last_time"`
}
