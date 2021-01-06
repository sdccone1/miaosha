package controller

import (
	"github.com/gin-gonic/gin"
	"miaosha/src/entity"
	"miaosha/src/service"
	"net/http"
)

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
		panic(err)
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
