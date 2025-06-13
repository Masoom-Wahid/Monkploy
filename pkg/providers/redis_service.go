package providers

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisService interface {
	GetRedisClient() *redis.Client
}

type redisService struct {
	redisClient *redis.Client
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Db       int
}

func NewRedisService(config RedisConfig) RedisService {
	return &redisService{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
			Password: config.Password,
			DB:       config.Db,
		}),
	}
}

func (rs *redisService) GetRedisClient() *redis.Client {
	return rs.redisClient
}
