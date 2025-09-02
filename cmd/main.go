package main

import (
	"context"
	_ "database/sql"
	"firstTestTask/internal/config"
	apphttp "firstTestTask/internal/delivery/http"
	"firstTestTask/internal/repository"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
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

	//роутинг
	router := apphttp.NewRouter(orderRepo)
	serv := &http.Server{Addr: ":8080", Handler: router}
	go func() {
		fmt.Println("HTTP ON 8080")
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//wg := sync.WaitGroup{}
	//wg.Add(1)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//go events.RunConsumer(ctx, "localhost:9092", orderRepo)
	//go func() {
	//	defer wg.Done()
	//	events.RunConsumer(ctx, "localhost:8080", orderRepo)
	//}()
	//wg.Wait()
	_ = serv.Shutdown(ctx)
	log.Println("bye")
}
