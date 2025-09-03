package kafka

import (
	"context"
	"encoding/json"
	"firstTestTask/internal/config"
	"firstTestTask/internal/domain"
	"firstTestTask/internal/repository"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	ord    *repository.OrderRepo
}

func NewConsumer(cfg *config.KafkaConfig, repo *repository.OrderRepo) *Consumer {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaBroker},
		Topic:   cfg.KafkaTopic,
		GroupID: "new-order",
	})
	return &Consumer{kafkaReader, repo}
}

func (cons *Consumer) Run(ctx context.Context) error {
	defer cons.reader.Close()
	for {
		msg, err := cons.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		var orderDto domain.OrderDTO
		if err := json.Unmarshal(msg.Value, &orderDto); err != nil {
			log.Println("bad msg: %v", err)
			continue
		}
		if err := cons.ord.SaveOrder(ctx, orderDto); err != nil {
			log.Println("repo error: %v", err)
		}
	}
}
