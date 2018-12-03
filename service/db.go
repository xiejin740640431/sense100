package service

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"sense100/config"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func ConnectDB() {
	var err error
	fmt.Println(config.MySql)
	db, err = gorm.Open("mysql", config.MySql)
	if err != nil {
		fmt.Println(err)
	}
	//设置闲置连接数
	db.DB().SetMaxIdleConns(10)
	//设置最大打开的连接数
	db.DB().SetMaxOpenConns(120)
	//开启日志
	db.LogMode(true)
}

func DisConnectDB() {
	if db != nil {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	}
}
