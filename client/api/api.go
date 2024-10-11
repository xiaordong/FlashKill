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
		resp.Response(ctx, resp.WithCode(401), resp.WithMsg(err.Error()))
		return
	}
	var token string
	token, err = service.BuyerRegister(u)
	if err != nil {
		resp.Response(ctx, resp.WithCode(402), resp.WithMsg(err.Error()))
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
		resp.Response(ctx, resp.WithCode(401), resp.WithMsg("error:"+err.Error()))
		return
	}
	var token string
	token, err = service.SellerRegister(s)
	if err != nil {
		resp.Response(ctx, resp.WithCode(402), resp.WithMsg("error:"+err.Error()))
		return
	}
	ctx.Header("SellerToken", token)
	resp.Response(ctx, resp.WithData(token))
}
func BuyerLogin(c context.Context, ctx *app.RequestContext) {
	var b model.Buyers
	b.Token = string(ctx.GetHeader("BuyerToken"))
	if b.Token == "" {
		resp.Response(ctx, resp.WithCode(403), resp.WithMsg("error"), resp.WithData("valid BuyerToken"))
		return
	}
	if err := ctx.BindJSON(&b); err != nil {
		resp.Response(ctx, resp.WithCode(401), resp.WithMsg("error:"+err.Error()))
		return
	}
	if err := service.BuyerLogin(b); err != nil {
		resp.Response(ctx, resp.WithCode(402), resp.WithMsg(err.Error()))
		return
	}
	resp.Response(ctx, resp.WithData(b))
}
func SellerLogin(c context.Context, ctx *app.RequestContext) {
	var s model.Sellers
	var err error
	s.Token = string(ctx.GetHeader("SellerToken"))
	if s.Token == "" {
		resp.Response(ctx, resp.WithCode(403), resp.WithMsg("error"), resp.WithData("valid SellerToken"))
		return
	}
	if err = ctx.BindJSON(&s); err != nil {
		resp.Response(ctx, resp.WithCode(401), resp.WithMsg("error:"+err.Error()))
		return
	}
	if _, err = service.SellerLogin(s); err != nil {
		resp.Response(ctx, resp.WithCode(402), resp.WithMsg(err.Error()))
		return
	}
	resp.Response(ctx, resp.WithData(s))
}
