package models

import (
	"awesomeProject/src/gin-blog/pkg/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {
	db, err := gorm.Open(
		settings.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			settings.DBUser,
			settings.DBPassWord,
			settings.DBHost,
			settings.DBName))
	if err != nil {
		log.Fatalf("init db error:%v\n", err)
		return
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return settings.DBTablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
func CloseDB() {
	defer db.Close()
}
