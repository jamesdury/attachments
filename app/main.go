package main

import (
	"log"
	"os"

	email "com.jamesdury.emailfiles/pkg/notmuch"
	"github.com/gofiber/fiber/v2"
	notmuch "github.com/zenhack/go.notmuch"
)

func main() {
	app := fiber.New()

	db := getDB()

	repo := email.NewRepo(db)
	service := email.NewService(repo)
	service.FetchEmail("tag:attachment")

	defer db.Close()

	log.Fatal(app.Listen(":8080"))
}

func getDB() *notmuch.DB {
	// TODO take cli argument or read env var
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	db, err := notmuch.Open(dirname+"/Mail", notmuch.DBReadOnly)
	if err != nil {
		log.Fatal("Database unavailable")
	}

	return db
}
