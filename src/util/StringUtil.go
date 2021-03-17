/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package util

import (
	"encoding/json"
	"go.uber.org/zap"
	"reflect"
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

func ObjectToString(obj interface{}) string {
	objectToString, err := json.Marshal(obj)
	if err != nil {
		zap.L().Error(err.Error())
	}
	return string(objectToString)
}

func StringToObject(str string, p reflect.Type) interface{} {
	obj := reflect.New(p)
	if err := json.Unmarshal([]byte(str), obj); err != nil {
		zap.L().Error(err.Error())
		return nil
	}
	return obj

}
