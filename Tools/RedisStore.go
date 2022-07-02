package Tools

import (
	"blog/model"
	"github.com/go-redis/redis/v8"
)

var RedisStore *redis.Client

func InitRedis(redisCfg model.RedisConfig) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Db,
	})
	RedisStore = client
}
