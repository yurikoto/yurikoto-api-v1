package service

import (
	"context"
	"yurikoto.com/yurikoto-api-go-v1/entity"
	"yurikoto.com/yurikoto-api-go-v1/redis"
	"yurikoto.com/yurikoto-api-go-v1/repository"
)

type SentenceService interface{
	Take() entity.Sentence
}

type sentenceService struct {
	sentenceRepository repository.SentenceRepository
}

func NewSentenceService(repo repository.SentenceRepository) SentenceService{
	return &sentenceService{
		sentenceRepository: repo,
	}
}

func (service *sentenceService) Take() entity.Sentence{
	rdb := redis.GetRedis()
	key := "sentence_requested"
	rdb.Incr(context.Background(), key)
	return service.sentenceRepository.Take()
}
