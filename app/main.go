package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"

	notmuch "github.com/zenhack/go.notmuch"

	home "github.com/jamesdury/attachments/app/home"

	email "github.com/jamesdury/attachments/pkg/notmuch"
)

//go:embed static/template/*
var embedDirTemplate embed.FS

//go:embed static/dist/*
var embedDirStatic embed.FS

func main() {
	engine := html.NewFileSystem(http.FS(embedDirTemplate), ".html")
	engine.AddFunc("filetype", TemplateFunctionFileType)
	engine.AddFunc("truncate", TemplateFunctionTruncate)
	engine.AddFunc("contact", TemplateFunctionContact)
	engine.AddFunc("email", TemplateFunctionEmail)
	engine.AddFunc("gravatar", TemplateFunctionGravatar)
	engine.AddFunc("prettydate", TemplateFunctionPrettyDate)
	engine.AddFunc("bytesize", TemplateFunctionBytesize)
	engine.AddFunc("escape", TemplateFunctionEscape)

	app := fiber.New(fiber.Config{
		Views: engine,
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
