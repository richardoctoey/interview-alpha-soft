package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db *gorm.DB
	once sync.Once
)

func init() {
	once.Do(func() {
		dsn := "root:gwudainget@tcp(127.0.0.1:3306)/dbmusic?charset=utf8mb4&parseTime=True&loc=Local"
		dbObj, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db = dbObj
	})
}

func GetDatabase() *gorm.DB {
	return db
}