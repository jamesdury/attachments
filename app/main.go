package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"

	home "github.com/jamesdury/attachments/app/home"
	email "github.com/jamesdury/attachments/pkg/notmuch"
	notmuch "github.com/zenhack/go.notmuch"
)

//go:embed static/template/*
var embedDirTemplate embed.FS

//go:embed static/dist/*
var embedDirStatic embed.FS

func main() {
	app := fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(embedDirTemplate), ".html"),
	})

	app.Use("/static", filesystem.New(filesystem.Config{
		Browse:     true,
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static/dist",
	}))

	db := getDB()

	emailService := email.NewService(email.NewRepo(db))

	home.Router(app.Group("/"), emailService)

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
