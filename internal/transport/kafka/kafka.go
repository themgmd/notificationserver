package kafka

import (
	"errors"
	"fmt"

	"github.com/onemgvv/notificationserver/internal/config"
	"github.com/onemgvv/notificationserver/internal/service"
	"github.com/segmentio/kafka-go"
)

type KafkaServer struct {
	reader              *kafka.Reader
	notificationService service.Notification
}

func New(notificationService service.Notification) (*KafkaServer, error) {
	cfg := config.Get()
	if cfg == nil {
		return nil, errors.New("config is not initialized")
	}

	kafkaCfg := cfg.Transport.Kafka
	addr := fmt.Sprintf("%s:%s", kafkaCfg.Host, kafkaCfg.Port)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{addr},
		Topic:    kafkaCfg.Topic,
		MaxBytes: kafkaCfg.MaxBytes,
	})

	return &KafkaServer{
		reader:              reader,
		notificationService: notificationService,
	}, nil
}
