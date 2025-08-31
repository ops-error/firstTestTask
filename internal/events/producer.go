package events

import (
	"context"
	"encoding/json"
	"firstTestTask/internal/models"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

type OrderEventProducer interface {
	Send(ctx context.Context, event models.OrderCreated) error
}

func NewProducer(address string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(address),
			Topic:    "orders.created",
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (prod *Producer) Send(ctx context.Context, event models.OrderCreated) error {
	bytes, _ := json.Marshal(event)
	return prod.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(event.OrderUID),
		Value: bytes,
	})
}
func (prod *Producer) Close() error {
	return prod.writer.Close()
}
