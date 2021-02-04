/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package entity

import "time"

type LoginUser struct {
	// 每个成员后面都有标签, 表示在解析对应数据格式时绑定到对方的字段名
	Mobile   string ` xml:"mobile" json:"mobile"  binding:"required"`
	PassWord string `xml:"password" json:"password"  binding:"required"`
}

/**
  这里注意一下，如果单独使用validator则这里的tag中应使用validate作为key而不是binding，但是只要使用了gin框架则进行参数校验时使用的关键字一定是binding，
当然两者可以进行集成，集成之后按照gin的方式在tag中使用binding关键字绑定我们利用第三方validator自定义的参数校验器
*/
type RegisterUser struct {
	Mobile   string ` xml:"mobile" json:"mobile"  binding:"ValidateFormat" `
	UserName string `xml:"username" json:"username"  binding:"NotNullAndAdmin"`
	PassWord string `xml:"password" json:"password"  binding:"required"`
}

type UserInfo struct {
	Mobile        string
	UserName      string
	PassWord      string
	RegisterDate  time.Time
	LastLoginDate time.Time
}
