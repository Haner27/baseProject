package datasource

import (
	"baseProject/config"
	"github.com/go-redis/redis"
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
