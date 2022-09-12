package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_HOST    string
	DB_PORT    string
	DB_USER    string
	DB_PASS    string
	DB_NAME    string
	JWT_SECRET string
}

func IntiDB() *gorm.DB {
	dsn := "root:N#@98wrft45@tcp(127.0.0.1:3306)/portal?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	return db
}
