package router

import (
	"client/api"
	"client/middleware"
	"client/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() {
	tb := middleware.NewTokenBucket(400000, 100000)
	h := server.Default(server.WithHostPorts(":8080"))
	h.GET("/Ping", middleware.RLMiddleware(tb), func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, "Pong")
	})

	b := h.Group("/buyer")
	{
		b.POST("/register", api.BuyerRegister)
		b.GET("/login", api.BuyerLogin)
		b.POST("/join", service.JoinActivity)
		b.POST("/pay", service.DealOrder)
	}

	s := h.Group("/seller")
	s.Use(middleware.RLMiddleware(tb))
	{
		s.POST("/register", api.SellerRegister)
		s.GET("/login", api.SellerLogin)
		s.POST("/new", service.Create) //创建活动
	}

	h.Spin()

}
