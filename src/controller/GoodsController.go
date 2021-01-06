package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadGoodsController(r *gin.Engine) {
	goodsGroup := r.Group("/goods")
	{
		goodsGroup.GET("/to_list", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "goods/goodslist.html", "goodslist page")
		})
	}
}
