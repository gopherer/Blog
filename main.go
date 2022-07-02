package main

import (
	"blog/Controller"
	"blog/Tools"
	"blog/model"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"os"
)

//执行顺序 全局变量→init方法→main方法
var cfg *model.Config

func init() {
	var err error
	cfg, err = Tools.ParseConfig("./Config/Blog.json") //cfg, err := Tools.ParseConfig("./Config/Blog.json") cfg未成功赋给全局变量cfg
	if err != nil {
		logger.Error("Tools.ParseConfig err")
		os.Exit(1)
	}
	err = Tools.OrmEngine(cfg.Database)
	if err != nil {
		logger.Error("Tools.OrmEngine err")
		os.Exit(1)
	}
	Tools.InitRedis(cfg.Redis)
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("Html/**/*")
	Tools.InitSession(engine, &cfg.Redis)
	userRouter := engine.Group("/user")
	new(Controller.UserController).UserController(userRouter)
	err := engine.Run(cfg.BlogHost + ":" + cfg.BlogPort)
	if err != nil {
		logger.Error(err)
	}
}
