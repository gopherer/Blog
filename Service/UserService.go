package Service

import (
	"blog/DataAccess"
	"blog/Model"
	"errors"
)

type UserService struct {
}

func (userService *UserService) UserLogin(reqUser Model.User) (*Model.User, error) {
	var userDao DataAccess.UserDao
	respUser := userDao.GetUserLogin(reqUser)
	if respUser.UserAccount == "" {
		return new(Model.User), errors.New("用户不存在")
	} else if respUser.UserAccount == respUser.UserAccount && reqUser.UserPwd != respUser.UserPwd {
		return new(Model.User), errors.New("输入密码有误")
	} else {
		return &respUser, nil
	}
}

func (userService *UserService) AccountModify(reqUser Model.User) error {

	//var userDao DataAccess.UserDao
	//userDao.ModifyAccount(reqUser)
	new(DataAccess.UserDao).ModifyAccount(reqUser)
	return nil
}

func (userService *UserService) ProfileModify(reqUser Model.User) error {
	new(DataAccess.UserDao).ModifyProfile(reqUser)
	return nil
}

func (userService *UserService) ProfileGet(reqUser Model.User) *Model.User {
	var userDao DataAccess.UserDao
	respUser := userDao.GetProfile(reqUser)
	return &respUser
}

func (userService *UserService) IconModify(reqUser Model.User) {
	new(DataAccess.UserDao).ModifyIcon(reqUser)
}

func (userService *UserService) GetUser() Model.User {
	return new(DataAccess.UserDao).GetUser()
}
