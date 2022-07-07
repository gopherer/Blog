package DataAccess

import (
	"blog/Model"
	"blog/Tools"
	_ "github.com/go-sql-driver/mysql" //未导入mysql驱动  会导致报错
	"github.com/wonderivan/logger"
)

type UserDao struct {
}

func (userDao *UserDao) GetUserLogin(reqUser Model.User) Model.User {
	var user Model.User
	has, err := Tools.DbEngine.Where("user_account=?", reqUser.UserAccount).Get(&user) //注意传入的是地址
	if !has {
		logger.Error(err)
	}
	return user
}

func (userDao *UserDao) ModifyAccount(reqUser Model.User) {
	_, _ = Tools.DbEngine.Cols("user_account", "user_pwd").Update(reqUser) //只更新数据表中user_account 和 user_pwd
}

func (userDao *UserDao) ModifyProfile(reqUser Model.User) {
	_, _ = Tools.DbEngine.Cols("user_nick", "user_profile", "user_contact").Update(reqUser)
}

func (userDao *UserDao) ModifyIcon(reqUser Model.User) {
	_, _ = Tools.DbEngine.Cols("user_icon").Update(reqUser)
}

func (userDao *UserDao) GetProfile(reqUser Model.User) Model.User {
	_, _ = Tools.DbEngine.Get(&reqUser) //注意传入的是地址
	return reqUser
}

func (userDao *UserDao) GetUser() Model.User {
	var user Model.User
	_, _ = Tools.DbEngine.Get(&user)
	return user
}
