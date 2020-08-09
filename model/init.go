package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

// DBEngin 是用来操作
var DBEngin *gorm.DB

func init() {
	var err error
	DBEngin, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/todolist?charset=utf8&parseTime=True&loc=Local")
	// DBEngin, err = gorm.Open("mysql", "root:nncz202008@tcp(sh-cdb-il38g3ny.sql.tencentcdb.com:60916)/todolist?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("连接数据库失败")
	} else {
		fmt.Println("连接数据库成功!")
	}

	DBEngin.AutoMigrate(&TodoModel{})
}
