/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"miaosha/src/dao"
	"miaosha/src/entity"
	"miaosha/src/redisService"
	"miaosha/src/util"
	"strconv"
	"time"
)

const (
	USERSIDEXPIRATION    time.Duration = time.Second * 3600 * 24 * 7
	USERCOOKIEEXPIRATION int           = 3600 * 24 * 7
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

func UserLogin(ctx *gin.Context, user *entity.LoginUser) int {
	usid := strconv.FormatInt(util.GetSnowflakeId(), 10)
	fmt.Printf("userSID= %v \n", usid)
	ustr := util.ObjectToString(user)
	fmt.Printf("ustr=%s \n", ustr)
	if sid := util.GetSidFromCookie(ctx); sid != "" {
		if !redisService.IsRedisDefaultClusterClientPoolAlive(ctx) {
			return -1 // -1表示服务器内部出错
		}
		if redisService.Set(ctx, usid, ustr, USERSIDEXPIRATION) {
			fmt.Printf("user %v 登录成功\n", user.Mobile)
			return 0 // 0表示success
		}
		return -1
	}
	u := dao.GetUserByMobile(user.Mobile)
	if u == nil || user.PassWord != u.PassWord {
		return 1 //1表示参数错误
	}
	if !redisService.IsRedisDefaultClusterClientPoolAlive(ctx) {
		return -1 // -1表示服务器内部出错
	}
	if redisService.Set(ctx, usid, ustr, USERSIDEXPIRATION) {
		ctx.SetCookie("userSID", usid, USERCOOKIEEXPIRATION, "/", "localhost", false, false)
		fmt.Printf("user %v 登录成功\n", user.Mobile)
		return 0 // 0表示success
	}
	return -1
}

func GetUserInfo(mobile string) *entity.UserInfo {
	u := dao.GetUserByMobile(mobile)
	return u
}
