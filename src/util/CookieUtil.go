/**
 * @Author: David Ma
 * @Date: 2021/2/5
 */
package util

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func checkError(err error) {
	zap.L().Error(err.Error())
}

func GetSidFromCookie(ctx *gin.Context) string {
	sidFromUrl := ctx.Query("usid")
	if sidFromUrl != "" {
		return sidFromUrl
	}
	sidFromHeader, err := ctx.Cookie("usid")
	if err != nil {
		checkError(err)
		return ""
	}
	return sidFromHeader
}
