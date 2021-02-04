/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package dao

import (
	"miaosha/src/entity"
	"time"
)

func UpdateUser(u *entity.RegisterUser) bool {
	db := GetMySqlPoll()
	sql := "update miaosha_user set nick_name = ?, password = ? where mobile = ?"
	stmt := GetPreparedStmt(db, sql)
	if _, err := stmt.Exec(u.UserName, u.PassWord, u.Mobile); err != nil {
		checkError(err)
	}
	defer stmt.Close()
	return true
}

func GetUserByMobile(mobile string) *entity.UserInfo {
	db := GetMySqlPoll()
	sql := "select distinct mobile,nick_name,password,register_date,last_login_date from miaosha_user where mobile = ?"
	stmt := GetPreparedStmt(db, sql)
	userInfo := new(entity.UserInfo)
	rows, err := stmt.Query(mobile)
	checkError(err)
	if rows.Next() {
		err = rows.Scan(&userInfo.Mobile, &userInfo.UserName, &userInfo.PassWord, &userInfo.RegisterDate, &userInfo.LastLoginDate)
		checkError(err)
		return userInfo
	}
	defer stmt.Close()
	return nil
}

func GetUserByName(name string) *entity.UserInfo {
	db := GetMySqlPoll()
	sql := "select distinct mobile,nick_name,password,register_date,last_login_date from miaosha_user where nick_name = ?"
	stmt := GetPreparedStmt(db, sql)
	userInfo := new(entity.UserInfo)
	rows, err := stmt.Query(name)
	checkError(err)
	if rows.Next() {
		err = rows.Scan(&userInfo.Mobile, &userInfo.UserName, &userInfo.PassWord, &userInfo.RegisterDate, &userInfo.LastLoginDate)
		checkError(err)
		return userInfo
	}
	defer stmt.Close()
	return nil
}

func InsertUser(user *entity.RegisterUser) bool {
	db := GetMySqlPoll()
	sql := "insert into miaosha_user(mobile,nick_name,password,register_date,last_login_date)values(?,?,?,?,?)"
	stmt := GetPreparedStmt(db, sql)
	_, err := stmt.Exec(user.Mobile, user.UserName, user.PassWord, time.Now(), time.Now())
	checkError(err)
	defer stmt.Close()
	return true
}
