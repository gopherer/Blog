package Controller

import (
	"blog/Controller/CTools"
	"blog/Service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (indexController *IndexController) IndexController(context *gin.Engine) {
	context.GET("/", HomePage)
}
func HomePage(context *gin.Context) {
	//user, _ := CTools.GetUserSession(context)
	user := new(Service.UserService).GetUser()
	blog := new(Service.BlogService).BlogAll()
	fmt.Println(user)
	fmt.Println(blog)
	CTools.ContextHtml(context, "Html/index.html", user, blog)
	//context.HTML(http.StatusOK, "Html/index.html", gin.H{
	//	"account": user.UserAccount,
	//	"nick":    user.UserNick,
	//	"icon":    user.UserIcon[2:],
	//	"profile": user.UserProfile,
	//	"contact": user.UserContact,
	//	"blog":    blog,
	//})
}
