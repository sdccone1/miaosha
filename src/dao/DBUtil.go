/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package dao

import (
	"database/sql"
	"go.uber.org/zap"

	//导入了但没有显示使用为什么还能通过sql.open("mysql")找到github.com/go-sql-driver/mysql这个第三方mysql驱动，是因为在其init()中github.com/go-sql-driver/mysql会向database/sql中注册一个mysql驱动
	//因为Go 语言对没有使用的导入是非常严格的。有时候程序员导入一个包可能只想要使用 init 函数的功能，例如一些引导工作。空白标示符就是一个不错的方式：比如下面这个
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func checkError(err error) {
	if err != nil {
		zap.L().Error(err.Error())
	}
}

func GetMySqlPoll() *sql.DB {
	if db != nil {
		return db
	}
	db, err := sql.Open("mysql", "root:davidmaqq0@tcp(127.0.0.1:3306)/miaoshadb?charset=utf8&parseTime=true&loc=Local") //3306是mysql的默认端口号
	checkError(err)
	err = db.Ping()
	checkError(err)
	//连接池中的最大链接数
	db.SetMaxOpenConns(100)
	//连接池中的最大空闲链接数
	db.SetMaxIdleConns(30)
	return db
}

func GetPreparedStmt(db *sql.DB, sql string) *sql.Stmt {
	stmt, err := db.Prepare(sql)
	checkError(err)
	return stmt
}
