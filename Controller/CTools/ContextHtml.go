package CTools

import (
	"blog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ContextHtml(context *gin.Context, user model.User, blog []model.Blog) {
	context.HTML(http.StatusOK, "home.html", gin.H{
		"account": user.UserAccount,
		"nick":    user.UserNick,
		"icon":    user.UserIcon[2:],
		"profile": user.UserProfile,
		"contact": user.UserContact,
		"blog":    blog,
	})
}
