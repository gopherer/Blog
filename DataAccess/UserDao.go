package DataAccess

import (
	"blog/Tools"
	"blog/model"
	_ "github.com/go-sql-driver/mysql" //未导入mysql驱动  会导致报错
	"github.com/wonderivan/logger"
)

type UserDao struct {
}

func (userDao *UserDao) GetUserInfo(reqUser model.User) model.User {
	var user model.User
	has, err := Tools.DbEngine.Where("user_account=?", reqUser.UserAccount).Get(&user) //注意传入的是地址
	if !has {
		logger.Error(err)
	}
	return user
}

func (userDao *UserDao) ModifyAccount(reqUser model.User) {
	_, _ = Tools.DbEngine.Cols("user_account", "user_pwd").Update(reqUser) //只更新数据表中user_account 和 user_pwd
}

func (userDao *UserDao) ModifyProfile(reqUser model.User) {
	_, _ = Tools.DbEngine.Cols("user_nick", "user_icon", "user_profile", "user_contact").Update(reqUser)
}

func (userDao *UserDao) GetProfile(reqUser model.User) model.User {
	_, _ = Tools.DbEngine.Get(&reqUser) //注意传入的是地址
	return reqUser
}
