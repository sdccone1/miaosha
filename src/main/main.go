/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package main

import (
	"go.uber.org/zap"
	"miaosha/src/controller"
)

func main() { //在运行时，不要直接右键run,因为此时会导致项目的根路径为这个main函数所在的路径，而不是真正的项目的根路径
	r := controller.SetUpRounter()
	controller.LoadUserController(r)
	controller.LoadGoodsController(r)
	err := r.Run("127.0.0.1:8080")
	if err != nil {
		zap.L().Panic(err.Error())
	}

}
