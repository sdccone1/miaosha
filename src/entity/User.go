package entity

import "time"

type LoginUser struct {
	// 每个成员后面都有标签, 表示在解析对应数据格式时绑定到对方的字段名
	Mobile   string ` xml:"mobile" json:"mobile"  binding:"required"`
	PassWord string `xml:"password" json:"password"  binding:"required"`
}

type RegisterUser struct {
	Mobile   string ` xml:"mobile" json:"mobile"  binding:"required"`
	UserName string `xml:"username" json:"username"  binding:"required"`
	PassWord string `xml:"password" json:"password"  binding:"required"`
}

type UserInfo struct {
	Mobile        string
	UserName      string
	PassWord      string
	RegisterDate  time.Time
	LastLoginDate time.Time
}
