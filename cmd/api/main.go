package main

import (
	"log/slog"
	"net/http"

	"github.com/alonsoF100/golos/internal/config"
	"github.com/alonsoF100/golos/internal/logger"
	"github.com/alonsoF100/golos/internal/repository/database/postgres"
	"github.com/alonsoF100/golos/internal/service"
	"github.com/alonsoF100/golos/internal/transport/http/handlers"
	"github.com/alonsoF100/golos/internal/transport/http/router"
	_ "github.com/alonsoF100/golos/migrations/postgres"
)

func main() {
	// Инициализация конфига
	config := config.Load()

	// Создание looger-а
	logger.Setup(config)

	// Создание pool-а
	pool, err := postgres.NewPool(config)
	if err != nil {
		slog.Error("Failed to create pool", "error", err)
	}
	defer pool.Close()
	slog.Info("Pool created successfully")

	// Создание слоя repo
	dataBase := postgres.New(pool)

	// Создание слоя service
	svc := service.New(
		dataBase, // user repo
		dataBase, // election repo
		dataBase, // voteVariat repo
		dataBase, // vote repo
	)

	// Создание слоя http
	handler := handlers.New(svc)

	// Сетап router-а
	router := router.New(handler).Setup()

	// Сетап сервера // TODO потом отдельный файл сделать с сетапом
	server := &http.Server{
		Addr:         config.Server.PortStr(),
		Handler:      router,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
	}

	// Запуск сервера
	slog.Info("Starting server", "port", config.Server.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed", "error", err)
	}
}
