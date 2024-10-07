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

	h.POST("/register", api.Register)

	h.Spin()
}
