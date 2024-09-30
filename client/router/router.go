package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() {
	h := server.Default()
	h.Group("/Ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, "Pong")
	})
	h.Spin()
}
