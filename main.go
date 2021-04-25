package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int64
	UserName string
	Email    string
}

func main() {
	//1. 我们在数据库操作的时候，比如dao层中当遇到一个sql.ErrNoRows的时候是否应该Wrap这个error抛给上层。

	//数据返回sql.ErrorNoRows不应该看作一个错误，我们应该允许返回空数据，进行降级处理不需要Wrap
	user, err := queryUserInfoByID(10)
	fmt.Printf("user:%v, err:%v\n", user, err)

}

func queryUserInfoByID(userId int64) (user *User, err error) {
	// 伪代码
	//err := queryUser()
	if err != nil && err == sql.ErrNoRows {
		err = nil
	}
	return
}
