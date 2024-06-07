package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/Morozov-rondo/restapi-template/config"
	"github.com/Morozov-rondo/restapi-template/internal/handlers"
	"github.com/Morozov-rondo/restapi-template/internal/repository"
	"github.com/Morozov-rondo/restapi-template/internal/router"
	"github.com/Morozov-rondo/restapi-template/internal/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Сборка и запуск приложения.
func Run() {

	// Инициализация конфигурации
	config.NewConfig()

	// Инициализация базы данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     config.Configs.DBHost,
		Port:     config.Configs.DBPort,
		Username: config.Configs.DBUsername,
		Password: config.Configs.DBPassword,
		DBName:   config.Configs.DBName,
	})

	if err != nil {
		log.Fatalf("Failed to initialize db: %s", err.Error())
	}

	// Инициализация мигратора
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// Инициализация мигратора
	m, err := migrate.New("file://./migrations", dsn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected successful")

	// Миграция схемы базы данных
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
	log.Println("Migration successful")

	// Инициализация сервисов
	repos := repository.NewRepository(db)
	s := service.NewService(repos)
	h := handlers.NewHandler(s)
	r := router.NewRouter(h)

	// Инициализация сервера
	server := http.Server{
		Addr:         fmt.Sprintf(":%s", config.Configs.ServerPort),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Запуск сервера
	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown при завершения работы сервера
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	err = db.Close()
	if err != nil {
		log.Fatalf("Database close error: %v", err)
	}
	<-ctx.Done()

	log.Println("Server stopped gracefully")

}
