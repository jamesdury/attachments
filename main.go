package main

import (
	"embed"
	"log"
	"os"

	notmuch "github.com/zenhack/go.notmuch"

	app "github.com/jamesdury/attachments/internal/server"
	email "github.com/jamesdury/attachments/pkg/notmuch"
	home "github.com/jamesdury/attachments/internal/home"
)

//go:embed static/template/*
var embedDirTemplate embed.FS

//go:embed static/dist/*
var embedDirStatic embed.FS

func main() {
	// TODO take cli argument or read env var
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	db, err := notmuch.Open(dirname+"/Mail", notmuch.DBReadOnly)
	if err != nil {
		log.Fatal("Database unavailable")
	}

	app := app.Setup(
		embedDirTemplate,
		embedDirStatic,
	)

	emailService := email.NewService(email.NewRepo(db))

	home.Router(app.Group("/"), emailService)

	defer db.Close()

	log.Fatal(app.Listen(":8080"))
}
