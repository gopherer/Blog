package Service

import (
	"blog/DataAccess"
	"blog/model"
	"errors"
)

type UserService struct {
}

func (userService *UserService) UserLogin(reqUser model.User) (*model.User, error) {
	var userDao DataAccess.UserDao
	respUser := userDao.GetUserInfo(reqUser)
	if respUser.UserAccount == "" {
		return new(model.User), errors.New("用户不存在")
	} else if respUser.UserAccount == respUser.UserAccount && reqUser.UserPwd != respUser.UserPwd {
		return new(model.User), errors.New("输入密码有误")
	} else {
		return &respUser, nil
	}
}

func (userService *UserService) AccountModify(reqUser model.User) error {
	//var userDao DataAccess.UserDao
	//userDao.ModifyAccount(reqUser)
	new(DataAccess.UserDao).ModifyAccount(reqUser)
	return nil
}

func (userService *UserService) ProfileModify(reqUser model.User) error {
	new(DataAccess.UserDao).ModifyProfile(reqUser)
	return nil
}

func (userService *UserService) ProfileGet(reqUser model.User) *model.User {
	var userDao DataAccess.UserDao
	respUser := userDao.GetProfile(reqUser)
	return &respUser
}
