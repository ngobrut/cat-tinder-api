package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/ngobrut/cat-tinder-api/config"
	"github.com/ngobrut/cat-tinder-api/internal/handler"
	"github.com/ngobrut/cat-tinder-api/internal/repository"
	"github.com/ngobrut/cat-tinder-api/internal/usecase"
	"github.com/ngobrut/cat-tinder-api/pkg/postgres"
)

func Exec() (*fiber.App, error) {
	cnf := config.New()

	db, err := postgres.NewDBClient(cnf)
	if err != nil {
		return nil, err
	}

	repo := repository.New(db)
	uc := usecase.New(cnf, db, repo)
	app := handler.InitHTTPHandler(cnf, uc)

	return app, nil
}

func main() {
	srv, err := Exec()
	if err != nil {
		log.Fatalf("[app-exec-failed] \n%v\n", err)
	}

	log.Printf("app running on :%d", 8080)

	// Listen from a different goroutine
	go func() {
		if err := srv.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	log.Println("gracefully shutting down...")
	_ = srv.Shutdown()

	log.Println("app shutdown.")
}
