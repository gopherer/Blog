package Middleware

import (
	"blog/Tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

//判断用户是否有权访问（是否登录） 中间件
func JudgeMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := Tools.GetSess(context, Tools.SessionKey)
		if user == nil {
			context.JSON(http.StatusOK, gin.H{
				"message": "请先登入才可以访问哦",
			})
			//abort（）顾名思义就是终止的意思，也就是说执行该函数，会终止后面所有的该请求下的函数。
			context.Abort()
		}
	}
}
