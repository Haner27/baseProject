package datasource

func CloseResource(redisDb *RedisDb) {
	redisDb.Close()
}
