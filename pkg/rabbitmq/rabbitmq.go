package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{Connection: conn, Channel: ch}, nil
}

func (r *RabbitMQ) DeclareQueue(queueName string) error {
	_, err := r.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	return err
}
