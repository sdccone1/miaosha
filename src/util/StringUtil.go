package util

import "strconv"

func Str2Int(str string) int {
	data, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	rInt := int(data)
	return rInt
}
