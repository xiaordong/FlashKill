package model

import (
	"github.com/shopspring/decimal"
)

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
	GoodsName  string          `gorm:"type:varchar(255);not null"`
	Price      decimal.Decimal `gorm:"type:decimal(20,8);not null"`
	Left       uint            `json:"left"`
	TimeOut    int             `json:"time_out"`
}

type Orders struct {
	OrderID    uint            `gorm:"primaryKey"`
	BuyerID    uint            `json:"buyer_id"`
	SellerID   uint            `json:"seller_id"`
	GoodsName  string          `json:"goodes_name"`
	GoodsPrice decimal.Decimal `json:"goods_price"`
	Status     bool            `json:"status"`
	TimeOut    int             `json:"time_out"`
}
