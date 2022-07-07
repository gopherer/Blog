package Controller

import (
	"blog/Controller/CTools"
	"blog/Controller/Middleware"
	"blog/Model"
	"blog/Service"
	"blog/Tools"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
}

//每调用一次handle 都可以视为调用一次goroutine
func (uL *UserController) UserController(context *gin.RouterGroup) {
	context.POST("/home", Middleware.VisitIP(), postUserLogin)
	context.POST("/account", Middleware.JudgeLogin(), postAccount)
	context.POST("/profile", Middleware.JudgeLogin(), postProfile)
	context.POST("/upload", Middleware.JudgeLogin(), postUpLoad)
	context.GET("/login", Middleware.VisitIP(), getUserLogin)
	context.GET("/account", Middleware.JudgeLogin(), getAccount)
	context.GET("/profile", Middleware.JudgeLogin(), getProfile)
	context.GET("/upload", Middleware.JudgeLogin(), getUpFile)
	context.GET("/flush", Middleware.JudgeLogin(), getCleanFlush)
}

func postUserLogin(context *gin.Context) {
	var user Model.User
	err := context.Bind(&user)
	if err != nil {
		logger.Error(err)
	}
	fmt.Println(user)
	respUser, err := new(Service.UserService).UserLogin(user)
	accountErr := errors.New("用户不存在")
	pwdErr := errors.New("输入密码有误")
	if err == nil {
		//context.JSON(http.StatusOK, gin.H{
		//		//	"mes":   "登陆成功",
		//		//	"hello": respUser,
		//		//})
		sess, _ := json.Marshal(respUser)
		//fmt.Println("111", string(sess))
		err := Tools.SetSess(context, Tools.SessionKey, sess)
		if err != nil {
			logger.Error(err)
			_, _ = context.Writer.WriteString("Session保存失败")
			os.Exit(1)
		}
		allBlog := new(Service.BlogService).BlogAll()
		//返回home结果页
		CTools.ContextHtml(context, *respUser, allBlog)
		//context.HTML(http.StatusOK, "home.html", gin.H{
		//	"account": respUser.UserAccount,
		//	"nick":    respUser.UserNick,
		//	"icon":    respUser.UserIcon[2:],
		//	"profile": respUser.UserProfile,
		//	"contact": respUser.UserContact,
		//})
	} else if err.Error() == accountErr.Error() {
		context.JSON(http.StatusOK, gin.H{
			"mes": "用户不存在",
		})
	} else if err.Error() == pwdErr.Error() {
		context.JSON(http.StatusOK, gin.H{
			"mes": "输入密码有误",
		})
	}
}

func postAccount(context *gin.Context) {
	var user Model.User
	//var blog Model.Blog
	err := context.Bind(&user)
	if err != nil {
		logger.Error(err)
	}
	err = new(Service.UserService).AccountModify(user)
	respUser := new(Service.UserService).ProfileGet(user)
	err = CTools.SetUserSession(context, *respUser)
	if err != nil {
		logger.Error(err)
	}
	//sess, _ := json.Marshal(respUser)
	//err = Tools.SetSess(context, Tools.SessionKey, sess)
	////fmt.Println("222", string(sess))
	//if err != nil {
	//	logger.Error(err)
	//	_, _ = context.Writer.WriteString("Session保存失败")
	//	os.Exit(1)
	//}
	//newBlog := new(Service.BlogService).BlogAll(blog)
	//fmt.Println(newBlog)
	allBlog := new(Service.BlogService).BlogAll()
	//返回home结果页
	CTools.ContextHtml(context, *respUser, allBlog)
	//context.HTML(http.StatusOK, "home.html", gin.H{
	//	"account": respUser.UserAccount,
	//	"nick":    respUser.UserNick,
	//	"icon":    respUser.UserIcon[2:],
	//	"profile": respUser.UserProfile,
	//	"contact": respUser.UserContact,
	//})

}

func postProfile(context *gin.Context) {
	var user Model.User
	//var sessionUser Model.User
	//var sessByte []byte
	err := context.Bind(&user)
	//fmt.Println(user)
	//user.UserNick = context.PostForm("user_nick")
	//file, err := context.FormFile("file")
	//user.UserProfile = context.PostForm("user_profile")
	//user.UserContact = context.PostForm("user_contact")
	if err != nil {
		logger.Error(err)
	}
	//fmt.Println(user)
	//fmt.Println(file)
	//fmt.Println(file.Filename)
	//if err != nil {
	//	context.String(500, "上传图片出错")
	//}
	//index := strings.Index(file.Filename, ".")
	//user.UserIcon = "./Upload/" + file.Filename[:index] + strconv.Itoa(int(time.Now().Unix())) + file.Filename[index:]
	err = new(Service.UserService).ProfileModify(user)
	if err != nil {
		logger.Error(err)
	}
	respUser := new(Service.UserService).ProfileGet(user)
	err = CTools.SetUserSession(context, *respUser)
	if err != nil {
		logger.Error(err)
	}
	//sess, _ := json.Marshal(respUser)
	//err = Tools.SetSess(context, Tools.SessionKey, sess)
	////fmt.Println("333", string(sess))
	//if err != nil {
	//	logger.Error(err)
	//	_, _ = context.Writer.WriteString("Session保存失败")
	//	os.Exit(1)
	//}
	//sess := CTools.GetSess(context, CTools.SessionKey)
	//switch sess.(type) {
	//case []byte:
	//	sessByte = sess.([]byte)
	//}
	//err = json.Unmarshal(sessByte, &sessionUser)
	//if err != nil {
	//	logger.Error(err)
	//}
	allBlog := new(Service.BlogService).BlogAll()
	//返回home结果页
	CTools.ContextHtml(context, *respUser, allBlog)
	//context.HTML(http.StatusOK, "home.html", gin.H{
	//	"account": respUser.UserAccount,
	//	"nick":    respUser.UserNick,
	//	"icon":    respUser.UserIcon[2:],
	//	"profile": respUser.UserProfile,
	//	"contact": respUser.UserContact,
	//})

}
func postUpLoad(context *gin.Context) {
	var user Model.User
	file, err := context.FormFile("file")
	if err != nil {
		context.String(500, "上传图片出错")
	}
	index := strings.Index(file.Filename, ".")
	//fmt.Println(file.Filename[:index])
	//fmt.Println(file.Filename[index:])
	filePath := "Upload/user/" + file.Filename[:index] + strconv.Itoa(int(time.Now().Unix())) + file.Filename[index:]
	err = context.SaveUploadedFile(file, filePath) //将接受的文件保存到filePath路径下
	if err != nil {
		logger.Error(err, "文件保存失败")
		return
	}
	user.UserIcon = filePath
	new(Service.UserService).IconModify(user) //将头像路径写入数据库
	respUser := new(Service.UserService).ProfileGet(user)
	err = CTools.SetUserSession(context, *respUser)
	if err != nil {
		logger.Error(err)
	}
	//sess, _ := json.Marshal(respUser)
	//err = Tools.SetSess(context, Tools.SessionKey, sess)
	//if err != nil {
	//	logger.Error(err)
	//	_, _ = context.Writer.WriteString("Session保存失败")
	//	os.Exit(1)
	//}
	allBlog := new(Service.BlogService).BlogAll()
	//返回home结果页
	CTools.ContextHtml(context, *respUser, allBlog)
	//context.HTML(http.StatusOK, "home.html", gin.H{
	//	"account": respUser.UserAccount,
	//	"nick":    respUser.UserNick,
	//	"icon":    respUser.UserIcon[2:],
	//	"profile": respUser.UserProfile,
	//	"contact": respUser.UserContact,
	//})

}

func getUserLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{})
}

func getAccount(context *gin.Context) {
	context.HTML(http.StatusOK, "account.html", gin.H{})
}

func getProfile(context *gin.Context) {

	context.HTML(http.StatusOK, "profile.html", gin.H{})
}

func getUpFile(context *gin.Context) {
	context.HTML(http.StatusOK, "upload.html", gin.H{})
}
func getCleanFlush(context *gin.Context) {
	Tools.RedisStore.FlushAll(context)
	context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}
