package entity

import "yurikoto.com/yurikoto-api-go-v1/config"

type Wallpaper struct {
	ID          uint64 `gorm:"primary_key;auto_increment;type:int(11)" json:"id"`
	Link        string `json:"link" binding:"required" gorm:"type:text"`
	Type        string `json:"type" binding:"required" gorm:"type:varchar(10)"`
	Orientation string `json:"orientation" binding:"required" gorm:"type:varchar(20)"`
}

func (Wallpaper) TableName() string {
	return config.Mysql.Prefix + "wallpapers"
}
