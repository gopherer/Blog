package main

import (
	"blog/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8")
	fmt.Println(err)
	if err := engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	var user model.User
	_, _ = engine.Where("user_account=?", "123456").Get(&user)
	fmt.Println(user)

}
