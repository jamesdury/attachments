package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"

	notmuch "github.com/zenhack/go.notmuch"

	"github.com/DusanKasan/parsemail"
)

func main() {
	readMail()
}

func readMail() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	db, err := notmuch.Open(dirname+"/Mail", notmuch.DBReadOnly)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	msgs, err := db.NewQuery("tag:attachment").Messages()
	if err != nil {
		panic(err)
	}
	var count int
	msg := &notmuch.Message{}
	for msgs.Next(&msg) {
		count++
		// invoke the GC to make sure it's running smoothly.
		if count%2 == 0 {
			runtime.GC()
		}
		readAttachments(msg.Filename())
	}
}

func readAttachments(name string) error {

	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	email, err := parsemail.Parse(scanner)
	if err != nil {
		return err
	}

	for _, a := range email.Attachments {
		fmt.Println(a.Filename)
	}

	return nil
}
