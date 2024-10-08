package resp

import (
	"github.com/cloudwego/hertz/pkg/app"
)

type Resp struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
	Data any    `json:"Data"`
}
type Option func(*Resp)

func WithMsg(msg string) Option {
	return func(r *Resp) { r.Msg = msg }
}
func WithCode(code int) Option {
	return func(r *Resp) { r.Code = code }
}
func WithData(data any) Option {
	return func(r *Resp) { r.Data = data }
}
func NewResponse(opt ...Option) *Resp {
	resp := &Resp{
		Code: 200,
		Msg:  "success",
		Data: nil,
	}
	for _, op := range opt {
		op(resp)
	}
	return resp
}
func Response(ctx *app.RequestContext, opt ...Option) {
	resp := NewResponse(opt...)
	ctx.JSON(resp.Code, resp)
}
