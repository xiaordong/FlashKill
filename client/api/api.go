package api

import (
	"client/model"
	"client/resp"
	"client/service"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

func Register(c context.Context, ctx *app.RequestContext) {
	fmt.Println("I'm in")
	var u model.Buyers
	var err error
	u.Name = ctx.PostForm("name")
	u.Password = ctx.PostForm("password")
	if err = ctx.Bind(&u); err != nil {
		resp.Response(ctx, resp.WithCode(401), resp.WithInfo("error"), resp.WithMsg(err.Error()))
		return
	}
	var token string
	if token, err = service.BuyerRegister(u); err != nil {
		resp.Response(ctx, resp.WithCode(402), resp.WithInfo(err.Error()))
		return
	}
	ctx.Header("userToken", token)
	resp.Response(ctx, resp.WithData(token))
}
