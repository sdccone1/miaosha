/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package service

import (
	"fmt"
	"go.uber.org/zap"
	"miaosha/src/dao"
	"miaosha/src/entity"
)

func UserUpdate(user *entity.RegisterUser) bool {
	return dao.UpdateUser(user)
}

func UserRegister(user *entity.RegisterUser) bool {
	u1 := dao.GetUserByMobile(user.Mobile)
	u2 := dao.GetUserByName(user.UserName)
	if u1 != nil || u2 != nil {
		return false
	}
	if dao.InsertUser(user) {
		zap.L().Info(fmt.Sprintf("user %v register successful", user))
		return true
	}
	return false
}

func GetUserInfo(mobile string) *entity.UserInfo {
	u := dao.GetUserByMobile(mobile)
	return u
}
