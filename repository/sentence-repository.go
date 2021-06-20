package repository

import (
	"gorm.io/gorm"
	"yurikoto.com/yurikoto-api-go-v1/entity"
)

type SentenceRepository interface {
	Take() entity.Sentence
	CloseDB()
}

type sentenceDatabase struct {
	connection *gorm.DB
}

func NewSentenceRepository() SentenceRepository {
	db := GetDB()
	err := db.AutoMigrate(&entity.Sentence{})
	if err != nil {
		panic(err.Error())
	}
	return &sentenceDatabase{
		connection: db,
	}
}

func (db *sentenceDatabase) CloseDB() {
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
 * @return entity.Sentence
 */
func (db *sentenceDatabase) Take() entity.Sentence {
	var sentence entity.Sentence
	db.connection.Set("gorm:auto_preload", true).Order("rand()").Take(&sentence)
	// db.connection.Set("gorm:auto_preload", true).Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)

	return sentence
}
