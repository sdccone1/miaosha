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
	"reflect"
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
	option := redisService.InitRedisClusterOption([]string{"192.168.211.129:7001", "192.168.211.129:7002", "192.168.211.129:7003"}, false,
		false, false, "davidmaqq0", 20, 5)
	redisClusterPool := redisService.GetRedisClusterConnectionPoll(ctx, option)
	if redisClusterPool == nil {
		return -1 // -1表示服务器内部出错
	}
	if sid := util.GetSidFromCookie(ctx); sid != "" {
		uFromRedis, err := redisClusterPool.Get(ctx, sid).Result()
		if err != nil {
			zap.L().Error(err.Error())
			return -1
		}
		obj := util.StringToObject(uFromRedis, reflect.TypeOf(user).Elem())
		if obj == nil {
			return -1
		}
		u := obj.(*entity.LoginUser)
		zap.L().Info(fmt.Sprintf("user: %v login success\n", u.Mobile))
		return 0
	}
	u := dao.GetUserByMobile(user.Mobile)
	if u == nil || user.PassWord != u.PassWord {
		return 1 //1表示参数错误
	}
	usid := strconv.FormatInt(util.GetSnowflakeId(), 10)
	ustr := util.ObjectToString(user)
	fmt.Printf("init userSID= %v ,ustr= %v \n", usid, ustr)
	if _, err := redisClusterPool.Set(ctx, usid, ustr, USERSIDEXPIRATION).Result(); err != nil {
		zap.L().Error(err.Error())
		return -1
	}
	zap.L().Info(fmt.Sprintf("Redis Set operation success,key = %s value = %s \n", usid, ustr))
	ctx.SetCookie("usid", usid, USERCOOKIEEXPIRATION, "/", "localhost", false, false)
	zap.L().Info(fmt.Sprintf("user %v 登录成功!\n", user.Mobile))
	return 0
}

func GetUserInfo(mobile string) *entity.UserInfo {
	u := dao.GetUserByMobile(mobile)
	return u
}
