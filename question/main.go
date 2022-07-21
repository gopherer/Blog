package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	//定义一个路由
	r := gin.Default()
	//加载静态文件
	r.Static("/dwz", "templates/statics")
	//添加自定义函数，解析下面的“百度一下”html代码
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//模板解析
	r.LoadHTMLGlob("templates/**/*")
	//模板渲染
	r.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='https://www.baidu.com'>百度一下</a>",
		})
	})
	r.Run(":9090")
}
