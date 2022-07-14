package CTools

import (
	"blog/Model"
	"blog/Service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ContextHtml(context *gin.Context, html string, user Model.User, blog []Model.Blog) {
	sort := new(Service.BlogService).GetSort()
	context.HTML(http.StatusOK, html, gin.H{
		"account": user.UserAccount,
		"nick":    user.UserNick,
		"icon":    user.UserIcon,
		"profile": user.UserProfile,
		"contact": user.UserContact,
		"blog":    blog,
		"sort":    sort,
	})
}
