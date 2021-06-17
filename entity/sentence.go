package entity

import "yurikoto.com/yurikoto-api-go-v1/config"

type Sentence struct{
	ID uint64 `gorm:"primary_key;auto_increment;type:int(11)" json:"id"`
	Content string `json:"content" binding:"required" gorm:"type:text"`
	Source string `json:"source" binding:"required" gorm:"type:varchar(256)"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Sentence) TableName() string {
	return config.Mysql.Prefix + "sentences"
}