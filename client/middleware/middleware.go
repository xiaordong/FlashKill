package middleware

import (
	"client/resp"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"sync"
	"time"
)

// TokenBucket 令牌桶算法
type TokenBucket struct {
	capacity int64 //最大容量，及分发的最大令牌数
	rate     int64 //每秒分发的速度
	tokens   int64
	mutex    sync.Mutex
}

func NewTokenBucket(capacity, rate int64) *TokenBucket {
	tb := &TokenBucket{
		capacity: capacity,
		rate:     rate,
		tokens:   capacity,
	}
	go tb.fill()
	return tb
}

func (tb *TokenBucket) fill() {
	lastTime := time.Now()
	for range time.Tick(time.Second) {
		tb.mutex.Lock()
		elapsed := time.Since(lastTime).Seconds()
		lastTime = time.Now()
		addedTokens := int64(elapsed * float64(tb.rate))
		if addedTokens > 0 {
			tb.tokens += addedTokens
			if tb.tokens > tb.capacity {
				tb.tokens = tb.capacity
			}
		}
		tb.mutex.Unlock()
	}
}

func (tb *TokenBucket) Take() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func RLMiddleware(tb *TokenBucket) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		if !tb.Take() {
			resp.Response(ctx, resp.WithCode(503), resp.WithMsg("too many requests"))
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
