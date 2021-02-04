/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package util

import (
	"go.uber.org/zap"
	"strconv"
)

func Str2Int(str string) int {
	data, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		//panic(err)
		zap.L().Error(err.Error())
	}
	rInt := int(data)
	return rInt
}
