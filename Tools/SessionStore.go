package Tools

import (
	"blog/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func InitSession(engine *gin.Engine, cfg *model.RedisConfig) {
	store, err := redis.NewStore(10, "tcp", cfg.Addr+":"+cfg.Port, "", []byte("secret"))
	if err != nil {
		logger.Error(err)
	}
	engine.Use(sessions.Sessions("mySession", store))
}

//set操作
func SetSess(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

//get操作
func GetSess(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}
