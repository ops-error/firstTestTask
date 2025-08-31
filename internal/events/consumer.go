package events

import (
	"context"
	"encoding/json"
	dto2 "firstTestTask/internal/models"
	"firstTestTask/internal/repository"
	"log"

	"github.com/segmentio/kafka-go"
)

func RunConsumer(ctx context.Context, address string, repo *repository.OrderRepo) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{address},
		Topic:       "orders.created",
		GroupID:     "order-worker",
		StartOffset: kafka.FirstOffset,
	})
	defer reader.Close()

	log.Println("Kafka consumer started")
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		var dto dto2.OrderDTO
		if err = json.Unmarshal(msg.Value, &dto); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}
		if err := repo.SaveOrder(ctx, dto); err != nil {
			log.Println("Error saving order:", err)
			continue
		}
		log.Println("Order saved", dto)
	}
}
