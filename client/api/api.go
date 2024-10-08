package api

import (
	"client/model"
	"client/resp"
	"client/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func BuyerRegister(c context.Context, ctx *app.RequestContext) {
	var u model.Buyers
	var err error
	u.Name = ctx.PostForm("name")
	u.Password = ctx.PostForm("password")
	if err = ctx.BindJSON(&u); err != nil {
		resp.Response(ctx, resp.WithCode(401), resp.WithInfo("error"), resp.WithMsg(err.Error()))
		return
	}
	var token string
	token, err = service.BuyerRegister(u)
	if err != nil {
		resp.Response(ctx, resp.WithCode(402), resp.WithInfo(err.Error()))
		return
	}
	ctx.Header("BuyerToken", token)
	resp.Response(ctx, resp.WithData(token))
}
func SellerRegister(c context.Context, ctx *app.RequestContext) {
	var s model.Sellers
	var err error
	s.Name = ctx.PostForm("name")
	s.Password = ctx.PostForm("password")
	if err = ctx.BindJSON(&s); err != nil {
		resp.Response(ctx, resp.WithCode(401), resp.WithInfo("error"), resp.WithMsg(err.Error()))
		return
	}
	var token string
	token, err = service.SellerRegister(s)
	if err != nil {
		resp.Response(ctx, resp.WithCode(402), resp.WithInfo(err.Error()))
		return
	}
	ctx.Header("SellerToken", token)
	resp.Response(ctx, resp.WithData(token))
}
