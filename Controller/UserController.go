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
)

type UserController struct {
}

//每调用一次handle 都可以视为调用一次goroutine
func (uL *UserController) UserController(context *gin.RouterGroup) {
	context.POST("/home", postUserLogin)
	context.POST("/account", Middleware.JudgeMiddle(), postAccount)
	context.POST("/profile", Middleware.JudgeMiddle(), postProfile)
	context.GET("/login", getUserLogin)
	context.GET("/account", Middleware.JudgeMiddle(), getAccount)
	context.GET("/profile", Middleware.JudgeMiddle(), getProfile)
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
			"icon":    respUser.UserIcon,
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
		"icon":    respUser.UserIcon,
		"profile": respUser.UserProfile,
		"contact": respUser.UserContact,
	})

}

func postProfile(context *gin.Context) {
	var user model.User
	//var sessionUser model.User
	//var sessByte []byte
	err := context.Bind(&user)
	if err != nil {
		logger.Error(err)
	}
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
		"icon":    respUser.UserIcon,
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
