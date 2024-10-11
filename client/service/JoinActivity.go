package service

import (
	"client/model"
	"client/resp"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"strconv"
	"time"
)

func JoinActivity(c context.Context, ctx *app.RequestContext) {
	var b model.Buyers
	var s model.Sellers
	b.Token = string(ctx.GetHeader("BuyerToken"))
	key := "lock"
	value := b.Token
	timeout := 10 * time.Second
	flag, err := model.RDB.SetNX(key, value, timeout).Result()
	if err != nil {
		panic(err)
	}
	if flag {
		fmt.Println("get the lock success")
		s.Token, err = model.RDB.Get("SellerToken").Result()
		if err != nil {
			log.Fatalln("get the SellerToken error:", err)
		}
		var left string
		left, err = model.RDB.Get(s.Token + "Left").Result()
		if err != nil {
			log.Fatalln("get the Seller Left error:", err)
		}
		s.Activity.Left, _ = strconv.Atoi(left)
		if s.Activity.Left > 0 {
			s.Activity.Left--
		} else {
			log.Println("seller left is 0")
			resp.Response(ctx, resp.WithMsg("seller left is 0"))
		}
		model.RDB.Set(s.Token+"Left", s.Activity.Left, 0)
		model.RDB.Set(b.Token+"Order", true, 30*time.Second)
		res, err := model.RDB.Get(key).Result()
		if err == nil && res == value {
			model.RDB.Del("lock")
		} else {
			log.Fatalln("Del lock error:", err)
		}
	} else {
		log.Println("get lock error")
		resp.Response(ctx, resp.WithCode(403), resp.WithMsg("too much request"))
	}

	resp.Response(ctx, resp.WithData("ordered successfully"))
}
func DealOrder(c context.Context, ctx *app.RequestContext) {
	var b model.Buyers
	var s model.Sellers
	b.Token = string(ctx.GetHeader("BuyerToken"))
	_, err := model.RDB.Get(b.Token + "Order").Result()
	if err != nil {
		s.Token, err = model.RDB.Get("SellerToken").Result()
		if err != nil {
			log.Fatalln("Get sellerToken error:", err)
		}
		model.RDB.Incr(s.Token + "Left")
		resp.Response(ctx, resp.WithData("order returned successfully"))
		return
	} else {
		resp.Response(ctx, resp.WithData("Order payed successfully"))
	}
}
