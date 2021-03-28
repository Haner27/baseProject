package datasource

import (
	"baseProject/config"
	"github.com/go-redis/redis"
	"time"
)

const (
	SlideWindowLimiterLuaScript = `
	local key = KEYS[1]
	local currentTS = tonumber(ARGV[1])
    local times = tonumber(ARGV[2])
	local intervals = tonumber(ARGV[3])
	
	local spanFirst = currentTS - intervals * 1000
	redis.call('ZREMRANGEBYSCORE', key, 0, spanFirst)
    redis.call('ZADD', key, currentTS, currentTS)
	if redis.call('ZCOUNT', key, spanFirst, currentTS) > times then
		return 0
	else
		return 1
	end
	`

	TokenBucketLimiterLuaScript = `
	local key = KEYS[1]
	local currentTimeStamp = tonumber(ARGV[1])
    local times = tonumber(ARGV[2])
	local intervals = tonumber(ARGV[3])
	local rate = times / (intervals * 1000)
	local lastTimeKey = 'lastTime'
	local currentTokenKey = 'currentToken'
	local lastTime = tonumber(redis.call('HGET', key, lastTimeKey))
	local currentToken = tonumber(redis.call('HGET', key, currentTokenKey))

	if lastTime then
	else
		lastTime = currentTimeStamp
	end

	if currentToken then
	else
		currentToken = times
	end

	local deltaToken = (currentTimeStamp - lastTime) * rate
	currentToken = currentToken + deltaToken
	if currentToken > times then
		currentToken = times
	end

	if currentToken - 1 > 0 then
		currentToken = currentToken - 1
		redis.call('HSET', key, lastTimeKey, currentTimeStamp)
		redis.call('HSET', key, currentTokenKey, currentToken)
		return 1
	else
		return 0
	end
	`
)

type RedisDb struct {
	Cli *redis.Client
}

func NewRedisDb(conf *config.Config) *RedisDb {
	options := redis.Options{
		Network:            "tcp",
		Addr:               conf.Redis.GetDSN(),
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
	db := redis.NewClient(&options)
	return &RedisDb{db}
}

func (r *RedisDb) Close() {
	r.Cli.Close()
}

func (r *RedisDb) SlideWindowLimiter(key string, times, intervals int) bool {
	currentTimeStamp := time.Now().UnixNano() / 1e6
	val, _ := r.Cli.Eval(SlideWindowLimiterLuaScript, []string{key}, currentTimeStamp, times, intervals).Result()
	if val.(int64) == 1 {
		return true
	}
	return false
}

func (r *RedisDb) TokenBucketLimiter(key string, times, intervals int) bool {
	currentTimeStamp := time.Now().UnixNano() / 1e6
	val, _ := r.Cli.Eval(TokenBucketLimiterLuaScript, []string{key}, currentTimeStamp, times, intervals).Result()
	if val.(int64) == 1 {
		return true
	}
	return false
}
