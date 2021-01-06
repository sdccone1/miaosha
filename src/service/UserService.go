package service

import (
	"miaosha/src/dao"
	"miaosha/src/entity"
)

func UserUpdate(user *entity.RegisterUser) bool {
	return dao.UpdateUser(user)
}

func UserRegister(user *entity.RegisterUser) bool {
	u := dao.GetUserByMobile(user.Mobile)
	if u != nil {
		return false
	}
	return dao.InsertUser(user)
}

func GetUserInfo(mobile string) *entity.UserInfo {
	u := dao.GetUserByMobile(mobile)
	return u
}
