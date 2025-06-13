package providers

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	VirtualHost string
}

type RabbitMQService interface {
	GetRabbitMQClient() *amqp.Connection
}

type rabbitMQService struct {
	connection *amqp.Connection
}

func NewRabbitMQService(config RabbitMQConfig) RabbitMQService {
	connURL := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.VirtualHost)
	conn, err := amqp.Dial(connURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}
	return &rabbitMQService{
		connection: conn,
	}
}

func (rs *rabbitMQService) GetRabbitMQClient() *amqp.Connection {
	return rs.connection
}
