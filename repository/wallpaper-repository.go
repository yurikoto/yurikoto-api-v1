package repository

import (
	"gorm.io/gorm"
	"yurikoto.com/yurikoto-api-go-v1/entity"
)

type WallpaperRepository interface {
	Take(t string) entity.Wallpaper
	CloseDB()
}

type wallpaperDatabase struct {
	connection *gorm.DB
}

func NewWallpaperRepository() *wallpaperDatabase {
	db := GetDB()
	err := db.AutoMigrate(&entity.Wallpaper{})
	if err != nil {
		panic(err.Error())
	}
	return &wallpaperDatabase{
		connection: db,
	}
}

func (db *wallpaperDatabase) CloseDB() {
	sqlDB, err := db.connection.DB()
	if sqlDB != nil {
		err = sqlDB.Close()
	}
	if err != nil {
		panic("Failed to close database")
	}
}

/**
 * @Description: 获取随机记录
 * @receiver db
 * @param t
 * @return entity.Wallpaper
 */
func (db *wallpaperDatabase) Take(t string) entity.Wallpaper {
	var wallpaper entity.Wallpaper
	db.connection.Set("gorm:auto_preload", true).Where("type = ?", t).Order("rand()").Take(&wallpaper)
	// db.connection.Set("gorm:auto_preload", true).Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)

	return wallpaper
}
