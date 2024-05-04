package main

import (
	"log"

	"github.com/ngobrut/cat-tinder-api/config"
	"github.com/ngobrut/cat-tinder-api/internal/handler"
	"github.com/ngobrut/cat-tinder-api/internal/repository"
	"github.com/ngobrut/cat-tinder-api/internal/usecase"
	"github.com/ngobrut/cat-tinder-api/pkg/postgres"
)

func Exec() error {
	cnf := config.New()

	db, err := postgres.NewDBClient(cnf)
	if err != nil {
		return err
	}

	repo := repository.New(db)
	uc := usecase.New(cnf, db, repo)
	app := handler.InitHTTPHandler(cnf, uc)

	log.Printf("app running on :%d", 8080)
	err = app.Listen(":8080")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Exec()
	if err != nil {
		log.Fatalf("[app-run-failed] \n%v\n", err)
	}
}
