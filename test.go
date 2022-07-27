package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name string
}

func main() {
	//t := time.Now()
	//a := t.Format("2006-01-02 15:04:05")
	//fmt.Println(a)
	var arr = make([]Person, 3)
	arr[0].Name = "0000"
	arr[1].Name = "1111"
	arr[2].Name = "2222"
	engine := gin.Default()
	engine.LoadHTMLFiles("test.html")
	engine.GET("/test", func(context *gin.Context) {
		context.HTML(http.StatusOK, "test.html", gin.H{
			"arr": arr,
		})
	})
	err := engine.Run()
	if err != nil {
		fmt.Println(err)
	}
}
