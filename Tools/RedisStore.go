package Tools

import (
	"blog/Model"
	"github.com/go-redis/redis/v8"
)

var RedisStore *redis.Client

func InitRedis(redisCfg Model.RedisConfig) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Db,
	})
	RedisStore = client
}
