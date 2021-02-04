/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"miaosha/src/entity"
	"miaosha/src/service"
	"net/http"
)

func checkError(err error) {
	zap.L().Error(err.Error())
}

func LoadUserController(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/to_register", toRegister)
		userGroup.POST("/do_register", doRegister)
	}
}

func userInfo(ctx *gin.Context) {
	mobile := ctx.Param("mobile")
	u := service.GetUserInfo(mobile)
	ctx.JSON(200, u)
}

func toRegister(ctx *gin.Context) {
	ctx.HTML(200, "user/register.html", nil)
}

func doRegister(ctx *gin.Context) {
	user := new(entity.RegisterUser)
	if err := ctx.ShouldBindJSON(user); err != nil {
		checkError(err)
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusOK, gin.H{ //踩坑，这里ctx.JSON()只是针对前端的request返回了一个body中封装了一个JSON的reponse，而不是意味着终止了整个方法的执行流程,要想手动终止方法的执行可以手动return！！！
				"status": 444,
				"msg":    "param is incorrect",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
		return
	}
	ok := service.UserRegister(user)
	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "注册成功，欢迎您：" + user.UserName,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "注册失败，用户已存在",
		})
	}
}
