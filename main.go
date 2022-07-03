package main

import (
	"blog/Controller"
	"blog/Tools"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"os"
)

//执行顺序 全局变量→init方法→main方法

func init() {
	var err error
	err = Tools.ParseConfig("./Config/Blog.json")
	if err != nil {
		logger.Error("Tools.ParseConfig err")
		os.Exit(1)
	}
	err = Tools.OrmEngine(Tools.Cfg.Database)
	if err != nil {
		logger.Error("Tools.OrmEngine err")
		os.Exit(1)
	}
	Tools.InitRedis(Tools.Cfg.Redis)
}

func main() {
	engine := gin.Default()
	engine.Static("user/Upload", "./Upload") //加载静态资源
	engine.LoadHTMLGlob("Html/**/*")
	//engine.MaxMultipartMemory = 8 << 20
	Tools.InitSession(engine, &Tools.Cfg.Redis)
	userRouter := engine.Group("/user")
	new(Controller.UserController).UserController(userRouter)
	err := engine.Run(Tools.Cfg.BlogHost + ":" + Tools.Cfg.BlogPort)
	if err != nil {
		logger.Error(err)
	}
}
