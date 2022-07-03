package Controller

import (
	"blog/Controller/Middleware"
	"blog/Service"
	"blog/Tools"
	"blog/model"
	"encoding/json"
	"errors"
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
	context.POST("/home", postUserLogin)
	context.POST("/account", Middleware.JudgeMiddle(), postAccount)
	context.POST("/profile", Middleware.JudgeMiddle(), postProfile)
	context.POST("/upLoad", Middleware.JudgeMiddle(), postUpLoad)
	context.GET("/login", getUserLogin)
	context.GET("/account", Middleware.JudgeMiddle(), getAccount)
	context.GET("/profile", Middleware.JudgeMiddle(), getProfile)
	context.GET("/upLoad", Middleware.JudgeMiddle(), getUpFile)
}

func postUserLogin(context *gin.Context) {
	var user model.User
	err := context.Bind(&user)
	if err != nil {
		logger.Error(err)
	}
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
		context.HTML(http.StatusOK, "home.html", gin.H{
			"account": respUser.UserAccount,
			"nick":    respUser.UserNick,
			"icon":    respUser.UserIcon[2:],
			"profile": respUser.UserProfile,
			"contact": respUser.UserContact,
		})
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
	var user model.User
	err := context.Bind(&user)
	if err != nil {
		logger.Error(err)
	}
	err = new(Service.UserService).AccountModify(user)
	respUser := new(Service.UserService).ProfileGet(user)
	sess, _ := json.Marshal(respUser)
	err = Tools.SetSess(context, Tools.SessionKey, sess)
	//fmt.Println("222", string(sess))
	if err != nil {
		logger.Error(err)
		_, _ = context.Writer.WriteString("Session保存失败")
		os.Exit(1)
	}
	context.HTML(http.StatusOK, "home.html", gin.H{
		"account": respUser.UserAccount,
		"nick":    respUser.UserNick,
		"icon":    respUser.UserIcon[2:],
		"profile": respUser.UserProfile,
		"contact": respUser.UserContact,
	})

}

func postProfile(context *gin.Context) {
	var user model.User
	//var sessionUser model.User
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
	sess, _ := json.Marshal(respUser)
	err = Tools.SetSess(context, Tools.SessionKey, sess)
	//fmt.Println("333", string(sess))
	if err != nil {
		logger.Error(err)
		_, _ = context.Writer.WriteString("Session保存失败")
		os.Exit(1)
	}
	//sess := Tools.GetSess(context, Tools.SessionKey)
	//switch sess.(type) {
	//case []byte:
	//	sessByte = sess.([]byte)
	//}
	//err = json.Unmarshal(sessByte, &sessionUser)
	//if err != nil {
	//	logger.Error(err)
	//}
	context.HTML(http.StatusOK, "home.html", gin.H{
		"account": respUser.UserAccount,
		"nick":    respUser.UserNick,
		"icon":    respUser.UserIcon[2:],
		"profile": respUser.UserProfile,
		"contact": respUser.UserContact,
	})

}
func postUpLoad(context *gin.Context) {
	var user model.User
	file, err := context.FormFile("file")
	if err != nil {
		context.String(500, "上传图片出错")
	}
	index := strings.Index(file.Filename, ".")
	//fmt.Println(file.Filename[:index])
	//fmt.Println(file.Filename[index:])
	filePath := "./Upload/" + file.Filename[:index] + strconv.Itoa(int(time.Now().Unix())) + file.Filename[index:]
	err = context.SaveUploadedFile(file, filePath) //将接受的文件保存到filePath路径下
	if err != nil {
		logger.Error(err, "文件保存失败")
		return
	}
	user.UserIcon = filePath
	new(Service.UserService).IconModify(user) //将头像路径写入数据库
	respUser := new(Service.UserService).ProfileGet(user)
	sess, _ := json.Marshal(respUser)
	err = Tools.SetSess(context, Tools.SessionKey, sess)
	if err != nil {
		logger.Error(err)
		_, _ = context.Writer.WriteString("Session保存失败")
		os.Exit(1)
	}
	context.HTML(http.StatusOK, "home.html", gin.H{
		"account": respUser.UserAccount,
		"nick":    respUser.UserNick,
		"icon":    respUser.UserIcon[2:],
		"profile": respUser.UserProfile,
		"contact": respUser.UserContact,
	})

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
