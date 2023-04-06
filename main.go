package main

import (
	"log"
	"os"

	notmuch "github.com/zenhack/go.notmuch"

	app "github.com/jamesdury/attachments/app"
)

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

	attachments := app.Build(db)

	defer db.Close()

	log.Fatal(attachments.Listen(":8080"))
}
