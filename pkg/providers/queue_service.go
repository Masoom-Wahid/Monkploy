package providers

import (
	"fmt"

	"github.com/hibiken/asynq"
)

type QueueService interface {
	GetQueueClient() *asynq.Client
}

type queueService struct {
	asynq *asynq.Client
}

func NewQueueService(config RedisConfig) QueueService {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.Db,
	})

	return &queueService{
		asynq: client,
	}
}

func (rs *queueService) GetQueueClient() *asynq.Client {
	return rs.asynq
}
