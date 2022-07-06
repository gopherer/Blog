package CTools

import (
	"blog/Tools"
	"blog/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func SetUserSession(context *gin.Context, respUser model.User) error {
	sess, _ := json.Marshal(respUser)
	err := Tools.SetSess(context, Tools.SessionKey, sess)
	//fmt.Println("333", string(sess))
	if err != nil {
		_, _ = context.Writer.WriteString("Session保存失败")
		return err
	}
	return nil
}

func GetUserSession(context *gin.Context) (model.User, error) {
	var sessUser model.User
	var sessByte []byte
	sess := Tools.GetSess(context, Tools.SessionKey)
	switch sess.(type) {
	case []byte:
		sessByte = sess.([]byte)
	}
	err := json.Unmarshal(sessByte, &sessUser)
	if err != nil {
		return *new(model.User), err
	}
	return sessUser, nil
}
