package model

type Sellers struct {
	SellerID uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"size:255;unique;not null"`
	Password string `json:"password" gorm:"size:255;not null"`
	Token    string `json:"token" gorm:"size:255;not null"`
}

type Buyers struct {
	BuyerID  uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Token    string `gorm:"type:varchar(255);not null"`
}
