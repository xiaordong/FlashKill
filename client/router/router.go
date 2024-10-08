package router

import (
	"client/api"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() {
	h := server.Default(server.WithHostPorts(":8080"))
	h.GET("/Ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, "Pong")
	})
	b := h.Group("/buyer")
	{
		b.POST("/register", api.BuyerRegister)
		b.GET("/login", api.BuyerLogin)
	}
	s := h.Group("/seller")
	{
		s.POST("/register", api.SellerRegister)
		s.GET("/login", api.BuyerLogin)
	}

	h.Spin()
}
