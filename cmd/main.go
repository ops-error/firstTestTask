package main

import (
	"context"
	_ "database/sql"
	"firstTestTask/internal/config"
	apphttp "firstTestTask/internal/delivery/http"
	"firstTestTask/internal/repository"
	"firstTestTask/internal/transport/kafka"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

//слои сервиса:
// Входящий слой: принять запрос,проверить авторизацию (при необходимости),парсинг JSON
// Контроллер: направить куда надо (маршрутизатор)
// Сервисы (бизнес-логика): Проверить данные, чтобы что-то от чего-то корректировалось и передать итог дальше
// Репозиторий: "Отнести и/или достать документы из PostgreSQL
// БД: Архив и картотека
//

func main() {
	newDB, err := config.NewDatabase()
	fmt.Println(newDB, err)
	defer func() {
		fmt.Println("Лавочка закрыта")
		newDB.Close()
	}()

	//producer := events.NewProducer("localhost:9092")
	//defer producer.Close()

	orderRepo := repository.NewOrderRepo(newDB)
	cfg := config.Load()
	cons := kafka.NewConsumer(cfg, orderRepo)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//cors
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	//роутинг
	router := apphttp.NewRouter(orderRepo)
	handler := corsConfig.Handler(router)
	serv := &http.Server{Addr: ":8080", Handler: handler}
	go func() {
		fmt.Println("HTTP ON 8080")
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("consumer started")
	if err := cons.Run(ctx); err != nil {
		log.Fatal("consumer: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = serv.Shutdown(ctx)
	log.Println("bye")
}
