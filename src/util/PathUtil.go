/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package util

import (
	"go.uber.org/zap"
	"os"
)

func GetRootDir() string {
	if dir, err := os.Getwd(); err != nil {
		zap.L().Error(err.Error())
	} else {
		return dir
	}
	return "/"
}
