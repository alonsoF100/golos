package main

import (
	"log"
	"net/http"

	"github.com/alonsoF100/golos/internal/repository/database/postgres"
	"github.com/alonsoF100/golos/internal/service"
	"github.com/alonsoF100/golos/internal/transport/http/handlers"
	"github.com/alonsoF100/golos/internal/transport/http/router"
)

const port = ":8080"
const connString = "взять из докера"

func main() {
	// Создание pool-а
	pool, err := postgres.NewPool(connString)
	if err != nil {
		log.Fatal("failed to pool")
	}
	defer pool.Close()

	// Создание слоя repo
	dataBase := postgres.New(pool)

	// Создание слоя service
	service := service.New(dataBase)

	// Создание слоя http
	handler := handlers.New(service)

	// Сетап router-а
	router := router.New(handler)

	// Запуск сервера
	http.ListenAndServe(port, router.Setup())
}

