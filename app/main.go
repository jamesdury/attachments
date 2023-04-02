package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	routes "github.com/jamesdury/attachments/app/routes"
	email "github.com/jamesdury/attachments/pkg/notmuch"
	notmuch "github.com/zenhack/go.notmuch"
)

//go:embed static/template/*
var embedDirTemplate embed.FS

func main() {
	app := fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(embedDirTemplate), ".html"),
	})

	db := getDB()

	emailRepo := email.NewRepo(db)
	emailService := email.NewService(emailRepo)

	home := app.Group("/")

	routes.HomeRouter(home, emailService)

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
