package rabbitmqserver

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	rabbitmqservice "github.com/whoami00911/gRPC-server/internal/service/rabbitmq"
	"github.com/whoami00911/gRPC-server/pkg/grpcPb"
	"github.com/whoami00911/gRPC-server/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
)

type rabbitmqServer struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	queue     amqp.Queue
	servicemq *rabbitmqservice.ServiceMq
	logger    *logger.Logger
}

func NewRabbitMQServer(servicemq *rabbitmqservice.ServiceMq, logger *logger.Logger) (*rabbitmqServer, error) {
	conn, err := amqp.Dial(viper.GetString("rabbitmq.url"))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := ch.QueueDeclare(
		viper.GetString("rabbitmq.queueName"), //имя очереди
		false,                                 // не durable
		false,                                 // удалять очередь, если не используется
		false,                                 // не эксклюзивная
		false,                                 // не ждать подтверждения
		nil,                                   // дополнительные аргументы
	)
	if err != nil {
		return nil, err
	}
	return &rabbitmqServer{
		conn:      conn,
		ch:        ch,
		queue:     queue,
		servicemq: servicemq,
		logger:    logger,
	}, nil
}

func (r *rabbitmqServer) ListenAndServeMq() error {
	logs, err := r.ch.Consume(
		viper.GetString("rabbitmq.queueName"),
		"",    // имя потребителя (оставляем пустым для автоматической генерации)
		true,  // автоматически подтверждать получение сообщения (auto-ack)
		false, // не эксклюзивно
		false, // не использовать no-local (не поддерживается RabbitMQ)
		false, // не ждать подтверждения
		nil,   // дополнительные аргументы
	)
	if err != nil {
		return err
	}
	wait := make(chan bool)

	for log := range logs {
		var logItem grpcPb.LogItem
		if err := bson.Unmarshal(log.Body, &logItem); err != nil {
			r.logger.Errorf("bson unmarshal error: %s", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := r.servicemq.Insert(ctx, &logItem); err != nil {
			r.logger.Errorf("Insert error: %s", err)
		}
		cancel()
	}
	<-wait
	return nil
}
