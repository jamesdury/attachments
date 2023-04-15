package main

import (
	"embed"
	"flag"
	"log"

	notmuch "github.com/zenhack/go.notmuch"

	home "github.com/jamesdury/attachments/internal/home"
	app "github.com/jamesdury/attachments/internal/server"
	email "github.com/jamesdury/attachments/pkg/notmuch"
)

//go:embed static/template/*
var embedDirTemplate embed.FS

//go:embed static/dist/*
var embedDirStatic embed.FS

func main() {
	dbPath := flag.String("db", "", "Provide path to the notmuch database directory (contains the .notmuch directory)")
	flag.Parse()

	if *dbPath == "" {
		log.Fatal("Database path not supplied")
	}

	db, err := notmuch.Open(*dbPath, notmuch.DBReadOnly)
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
