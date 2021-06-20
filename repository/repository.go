package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"yurikoto.com/yurikoto-api-go-v1/config"
)

var db *gorm.DB

/**
 * @Description: 静态mysql链接
 * @return *gorm.DB
 */
func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	dsn := config.Mysql.Username + ":" + config.Mysql.Pwd + "@tcp(" + config.Mysql.Host + ":" + config.Mysql.Port + ")/" + config.Mysql.Dbname + "?charset=" + config.Mysql.Charset + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
