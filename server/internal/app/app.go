package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cerque1/tutor-website-001/internal/config"
	"github.com/cerque1/tutor-website-001/internal/database"
	"github.com/cerque1/tutor-website-001/internal/handler"
	"github.com/cerque1/tutor-website-001/internal/migrator"
	"github.com/cerque1/tutor-website-001/internal/repository"
	"github.com/cerque1/tutor-website-001/internal/router"
	"github.com/cerque1/tutor-website-001/internal/service"
)

const filePath = "file://migrations"

type App struct {
    Server *http.Server
	DB *sql.DB
}

func New(cfg *config.Config) (*App, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	migrator.Apply(filePath, connStr)

	db, err := database.New(connStr)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	userRouter := router.New(router.Handlers{
		User: userHandler,
	})

	server := &http.Server{
		Addr: cfg.HTTPAddr,
		Handler: userRouter,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout: time.Minute,
	}
	
	return &App{
		Server: server,
		DB: db,
	}, nil
}

func Run() {
	cfg := config.Load()

	application, err := New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Printf("Сервер запущен на %s", application.Server.Addr)

		if err := application.Server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
				log.Fatal(err)
			}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Выключение")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := application.Server.Shutdown(ctx); err != nil {
		log.Printf("Ошибка завершения работы: %v", err)
	}

	if err := application.DB.Close(); err != nil {
		log.Printf("Ошибка базы данных: %v", err)
	}

	log.Println("Приложение остановлено")
}
