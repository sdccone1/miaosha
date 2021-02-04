/**
 * @author:David Ma
 * @date:2021-02-01
 */

package controller

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"miaosha/src/logService"
	"miaosha/src/util"
	"miaosha/src/validatorService"
	"time"
)

func SetUpRounter() *gin.Engine {
	r := gin.New()
	filePath := util.GetRootDir() + "/userLocal/test.log"
	//if _,err := os.Create(filePath); err != nil{
	//	zap.L().Error(err.Error())
	//}
	logger := logService.NewLogger(filePath, zapcore.InfoLevel, 64, 3, 7, true, "MiaoSha")
	// 将自定义的logger替换为全局的logger
	zap.ReplaceGlobals(logger)

	// 使用zap日志库
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	//使用自定义参数校验器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NotNullAndAdmin", validatorService.NameNotNullAndAdmin)
		v.RegisterValidation("ValidateFormat", validatorService.MobileFormatIsCorrect)
	}

	//告诉router要读取的模板文件的位置,都是从项目的根路径开始,如果templates文件夹下有多级目录(这里是二级目录,且用**表示模糊匹配)，则必须修改pattern参数来匹配：
	//且在对应的handlerfunction中，对应跳转的html的文件名也需要给指定路径(比如：user/login)，而且在对应的html模板中也需要相应处理
	r.LoadHTMLGlob("templates/**/*")
	//指定渲染页面时所需要使用到的静态文件，
	//@parm:relativePath:表示在html中所有请求路径为/static(也就是这里的/static跟/user类似，相当于一个URI)下的资源均从@parm:root所表示的项目路径下来找
	//且一般保持/static和static文件夹名一致，方便找，也可以不一致
	r.Static("/static", "static")
	return r
}
