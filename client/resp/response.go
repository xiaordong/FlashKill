package resp

import "github.com/cloudwego/hertz/pkg/app"

type Resp struct {
	code int
	info string
	msg  string
	data interface{}
}
type Option func(*Resp)

func WithMsg(msg string) Option {
	return func(r *Resp) { r.msg = msg }
}
func WithCode(code int) Option {
	return func(r *Resp) { r.code = code }
}
func WithInfo(info string) Option {
	return func(r *Resp) { r.info = info }
}
func WithData(data interface{}) Option {
	return func(r *Resp) { r.data = data }
}
func NewResponse(opt ...Option) *Resp {
	resp := &Resp{
		code: 200,
		msg:  "success",
		data: nil,
	}
	for _, op := range opt {
		op(resp)
	}
	return resp
}
func Response(ctx *app.RequestContext, opt ...Option) {
	resp := NewResponse(opt...)
	ctx.JSON(resp.code, resp)
}
