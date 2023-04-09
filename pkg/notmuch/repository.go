package email

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"

	//"net/http"
	"os"

	"github.com/DusanKasan/parsemail"
	notmuch "github.com/zenhack/go.notmuch"
)

type Repository interface {
	Query(query string) (*[]Email, error)
}

type repository struct {
	connection *notmuch.DB
}

func NewRepo(db *notmuch.DB) Repository {
	return &repository{
		connection: db,
	}
}

func readEmail(name string) (parsemail.Email, error) {
	file, err := os.Open(name)
	if err != nil {
		return parsemail.Email{}, nil
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	email, err := parsemail.Parse(scanner)
	if err != nil {
		return parsemail.Email{}, nil
	}

	return email, nil
}

// TODO tidy this function up
func (r *repository) Query(query string) (*[]Email, error) {
	q := r.connection.NewQuery(query)
	q.AddTagExclude("spam")
	q.SetSortScheme(notmuch.SORT_NEWEST_FIRST)
	msgs, err := q.Messages()

	if err != nil {
		return nil, err
	}
	msg := &notmuch.Message{}

	var emails []Email
	for msgs.Next(&msg) {
		email, err := readEmail(msg.Filename())
		if err == nil {
			for _, a := range email.Attachments {

				var bodyBytes []byte
				bodyBytes, err = io.ReadAll(a.Data)
				if err != nil {
					fmt.Printf("error: %v", err)
				}
				a.Data = io.NopCloser(bytes.NewBuffer(bodyBytes))
				b64 := base64.StdEncoding.EncodeToString(bodyBytes)

				// TODO Determine the content type of the image file
				// mimeType := http.DetectContentType(bodyBytes)
				// And if it matches what the email is saying
				// a.ContentType === http.DetectContentType

				e := Email{
					ContentType: a.ContentType,
					Data:        a.Data,
					Filename:    a.Filename,
					From:        msg.Header("From"),
					Received:    msg.Date(),
					Size:        filesize(b64),
					Subject:     email.Subject,
					ThreadId:    msg.ThreadID(),
				}

				emails = append(emails, e)
			}
		}
	}

	return &emails, nil
}
