package Controller

import (
	"blog/Controller/CTools"
	"blog/Controller/Middleware"
	"blog/Model"
	"blog/Service"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BlogController struct {
}

func (blogController *BlogController) BlogController(context *gin.RouterGroup) {
	context.POST("/publish", Middleware.JudgeLogin(), postPublish)
	context.POST("/modify", Middleware.JudgeLogin(), postModify)
	context.POST("/delete", Middleware.JudgeLogin(), postDelete)
	context.GET("/publish", Middleware.JudgeLogin(), getPublish)
	context.GET("/modify", Middleware.JudgeLogin(), getModify)
	context.GET("/delete", Middleware.JudgeLogin(), getDelete)
}
func postPublish(context *gin.Context) {
	var blog Model.Blog
	//var sessUser Model.User
	//var sessByte []byte
	file, err := context.FormFile("blog_photo")
	if err != nil {
		context.String(500, "上传图片出错")
	}
	index := strings.Index(file.Filename, ".")
	filePath := "Upload/blog/" + file.Filename[:index] + strconv.Itoa(int(time.Now().Unix())) + file.Filename[index:]
	err = context.SaveUploadedFile(file, filePath) //将接受的文件保存到filePath路径下

	blog.BlogTitle = context.PostForm("blog_title")
	blog.BlogPhoto = filePath
	blog.BlogContent = context.PostForm("blog_content")

	//获取系统时间 设置博客创建时间
	t := time.Now()
	blog.BlogCreateTime = t.Format("2006-01-02 15:04:05")
	blog.BlogModifyTime = t.Format("2006-01-02 15:04:05")

	err = new(Service.BlogService).PostPublish(blog)
	if err != nil {
		logger.Error(err)
	}
	//从session中获取model.user
	//sess := Tools.GetSess(context, Tools.SessionKey)
	//switch sess.(type) {
	//case []byte:
	//	sessByte = sess.([]byte)
	//}
	//err = json.Unmarshal(sessByte, &sessUser)
	//if err != nil {
	//	logger.Error(err)
	//}
	sessUser, err := CTools.GetUserSession(context)
	if err != nil {
		logger.Error(err)
	}
	//fmt.Println("sess", sessUser)
	allBlog := new(Service.BlogService).BlogAll()
	//返回home结果页
	CTools.ContextHtml(context, sessUser, allBlog)
	//context.HTML(http.StatusOK, "home.html", gin.H{
	//	//	"account": sessUser.UserAccount,
	//	//	"nick":    sessUser.UserNick,
	//	//	"icon":    sessUser.UserIcon[2:],
	//	//	"profile": sessUser.UserProfile,
	//	//	"contact": sessUser.UserContact,
	//	//	"title":   blog.BlogTitle,
	//	//})
}

//修改发布的博客 本质上就是种另类的发布博客    目前先这样处理
func postModify(context *gin.Context) {
	var blog Model.Blog
	//var sessUser Model.User
	//var sessByte []byte
	file, err := context.FormFile("blog_photo")
	if err != nil {
		context.String(500, "上传图片出错")
	}
	index := strings.Index(file.Filename, ".")
	filePath := "./Upload/blog/" + file.Filename[:index] + strconv.Itoa(int(time.Now().Unix())) + file.Filename[index:]
	err = context.SaveUploadedFile(file, filePath) //将接受的文件保存到filePath路径下

	blog.BlogTitle = context.PostForm("blog_title")
	blog.BlogPhoto = filePath
	blog.BlogContent = context.PostForm("blog_content")

	//修改点  ---  postPublish
	t := time.Now()
	blog.BlogModifyTime = t.Format("2006-01-02 15:04:05")
	err = new(Service.BlogService).PostModify(blog)
	//修改点  ---  postPublish
	if err != nil {
		logger.Error(err)
	}
	//从session中获取model.user
	//sess := Tools.GetSess(context, Tools.SessionKey)
	//switch sess.(type) {
	//case []byte:
	//	sessByte = sess.([]byte)
	//}
	//err = json.Unmarshal(sessByte, &sessUser)
	//if err != nil {
	//	logger.Error(err)
	//}
	sessUser, err := CTools.GetUserSession(context)
	if err != nil {
		logger.Error(err)
	}
	allBlog := new(Service.BlogService).BlogAll()
	//返回home结果页
	CTools.ContextHtml(context, sessUser, allBlog)
	//fmt.Println("sess", sessUser)
	//context.HTML(http.StatusOK, "home.html", gin.H{
	//	"account": sessUser.UserAccount,
	//	"nick":    sessUser.UserNick,
	//	"icon":    sessUser.UserIcon[2:],
	//	"profile": sessUser.UserProfile,
	//	"contact": sessUser.UserContact,
	//	"title":   blog.BlogTitle,
	//})
}
func postDelete(context *gin.Context) {
	var blog Model.Blog
	//var sessUser Model.User
	//var sessByte []byte
	//var sessUser Model.User
	//var sessByte []byte
	blog.BlogTitle = context.PostForm("blog_title")
	_ = new(Service.BlogService).PostDelete(blog)

	//sess := Tools.GetSess(context, Tools.SessionKey)
	//switch sess.(type) {
	//case []byte:
	//	sessByte = sess.([]byte)
	//}
	//err := json.Unmarshal(sessByte, &sessUser)
	//if err != nil {
	//	logger.Error(err)
	//}
	sessUser, err := CTools.GetUserSession(context)
	if err != nil {
		logger.Error(err)
	}
	allBlog := new(Service.BlogService).BlogAll()
	//返回home结果页
	CTools.ContextHtml(context, sessUser, allBlog)
	//从session中获取model.user
	//sess := CTools.GetSess(context, CTools.SessionKey)
	//switch sess.(type) {
	//case []byte:
	//	sessByte = sess.([]byte)
	//}
	//err := json.Unmarshal(sessByte, &sessUser)
	//if err != nil {
	//	logger.Error(err)
	//}
	//if err != nil {
	//	logger.Error(err)
	//}
	//
	////fmt.Println("sess", sessUser)
	//context.HTML(http.StatusOK, "home.html", gin.H{
	//	"account": sessUser.UserAccount,
	//	"nick":    sessUser.UserNick,
	//	"icon":    sessUser.UserIcon[2:],
	//	"profile": sessUser.UserProfile,
	//	"contact": sessUser.UserContact,
	//	"title":   blog.BlogTitle,
	//})
}

//使用路由组会导致访问时 默认会加上路由组定义的路径 导致页面跳转问题
func getPublish(context *gin.Context) {
	context.HTML(http.StatusOK, "publish.html", gin.H{})
}
func getModify(context *gin.Context) {
	context.HTML(http.StatusOK, "modify.html", gin.H{})
}
func getDelete(context *gin.Context) {
	context.HTML(http.StatusOK, "delete.html", gin.H{})
}
