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
	var a model.Activities
	s.Token = string(ctx.GetHeader("SellerToken"))
	if s.Token == "" {
		resp.Response(ctx, resp.WithCode(403), resp.WithMsg("error"), resp.WithData("valid SellerToken"))
		return
	}
	if err := ctx.BindJSON(&a); err != nil {
		resp.Response(ctx, resp.WithCode(401), resp.WithMsg("error:"+err.Error()))
		return
	}
	model.RDB.Set(s.Token+"GoodsName", a.GoodsName, time.Minute*time.Duration(a.TimeOut))
	model.RDB.Set(s.Token+"Price", a.Price, time.Minute*time.Duration(a.TimeOut))
	model.RDB.Set(s.Token+"Left", a.Left, time.Minute*time.Duration(a.TimeOut))
	resp.Response(ctx, resp.WithData(a))
}
