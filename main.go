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
		logger.Error("CTools.ParseConfig err")
		os.Exit(1)
	}
	err = Tools.OrmEngine(Tools.Cfg.Database)
	if err != nil {
		logger.Error("CTools.OrmEngine err")
		os.Exit(1)
	}
	Tools.InitRedis(Tools.Cfg.Redis)
}

func main() {
	engine := gin.Default()
	engine.Static("user/Upload/user", "./Upload/user") //加载静态资源   路由组会导致访问是默认加上路由组定义的路径。user组
	engine.Static("blog/Upload/user", "./Upload/user")
	//engine.Static("/blog/publish", "./Upload/blog")
	engine.Static("user/Upload/blog", "./Upload/blog") //加载静态资源   路由组会导致访问是默认加上路由组定义的路径。blog组
	engine.Static("blog/Upload/blog", "./Upload/blog")
	//
	engine.Static("Upload/user", "./Upload/user")
	engine.Static("Upload/blog", "./Upload/blog")
	engine.LoadHTMLGlob("Html/**/*") //模板解析
	//engine.MaxMultipartMemory = 8 << 20
	Tools.InitSession(engine, &Tools.Cfg.Redis)
	userRouter := engine.Group("/user")
	blogRouter := engine.Group("/blog")
	new(Controller.UserController).UserController(userRouter)
	new(Controller.BlogController).BlogController(blogRouter)
	new(Controller.IndexController).IndexController(engine)
	err := engine.Run(Tools.Cfg.BlogHost + ":" + Tools.Cfg.BlogPort)
	if err != nil {
		logger.Error(err)
	}
}
