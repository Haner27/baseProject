package middleware

import (
	"baseProject/datasource"
	"baseProject/util/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SliderWindowLimiter struct {
	cache *datasource.RedisDb
}

func NewSliderWindowLimiter(cache *datasource.RedisDb) *SliderWindowLimiter {
	return &SliderWindowLimiter{
		cache,
	}
}

func (s *SliderWindowLimiter) Limiter(times, intervals int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := fmt.Sprintf("SliderWindowLimiter:%d:%d:%s", intervals, times, ctx.ClientIP())
		if !s.cache.SlideWindowLimiter(key, times, intervals) {
			response.FrequentlyRequestResp(ctx)
			return
		}
		ctx.Next()
	}
}

type TokenBucketLimiter struct {
	cache *datasource.RedisDb
}

func NewTokenBucketLimiter(cache *datasource.RedisDb) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		cache,
	}
}

func (t *TokenBucketLimiter) Limiter(times, intervals int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := fmt.Sprintf("TokenBucketLimiter:%d:%d:%s", intervals, times, ctx.ClientIP())
		if !t.cache.TokenBucketLimiter(key, times, intervals) {
			response.FrequentlyRequestResp(ctx)
			return
		}
		ctx.Next()
	}
}
