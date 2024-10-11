package service

import (
	"client/model"
	"client/resp"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

func Create(c context.Context, ctx *app.RequestContext) {
	var s model.Sellers
	s.Token = string(ctx.GetHeader("SellerToken"))
	if s.Token == "" {
		resp.Response(ctx, resp.WithCode(403), resp.WithMsg("error"), resp.WithData("valid SellerToken"))
		return
	}
	if err := ctx.BindJSON(&s.Activity); err != nil {
		resp.Response(ctx, resp.WithCode(401), resp.WithMsg("error:"+err.Error()))
		return
	}
	model.RDB.Set(s.Token+"GoodsName", s.Activity.GoodsName, time.Minute*time.Duration(s.Activity.TimeOut))
	model.RDB.Set(s.Token+"Price", s.Activity.Price, time.Minute*time.Duration(s.Activity.TimeOut))
	model.RDB.Set(s.Token+"Left", s.Activity.Left, time.Minute*time.Duration(s.Activity.TimeOut))
	resp.Response(ctx, resp.WithData(s))
}
