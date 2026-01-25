package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rent/internal/config"
)

var DB *gorm.DB

func InitMySQL() *gorm.DB {
	dsn := config.Cfg.Mysql.Source
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	return DB
}
