package model

import (
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
)

var RDB *redis.Client

type Sellers struct {
	Name     string     `json:"name"`
	Password string     `json:"password"`
	Token    string     `json:"token"`
	Activity Activities `json:"activity"`
}

type Buyers struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Order    Orders `json:"order"`
}

type Activities struct {
	GoodsName string          `json:"goods_name"`
	Price     decimal.Decimal `json:"price" `
	Left      int             `json:"left"`
	TimeOut   int             `json:"time_out"`
}

type Orders struct {
	Status  bool `json:"status"`
	TimeOut int  `json:"time_out"`
}
