package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Static("/upload", "./Upload")
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("./test.tmpl")

	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.tmpl", `<div> <img src ="upload/go1656829298.jpg" /></div>`)
	})

	router.Run(":8080")
}
