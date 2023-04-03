package email

import (
	"bufio"
	"os"

	"github.com/DusanKasan/parsemail"
	notmuch "github.com/zenhack/go.notmuch"
)

type Repository interface {
	Fetch(query string) ([]Email, error)
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

func (r *repository) Fetch(query string) ([]Email, error) {
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

				/*
				 * TODO write a check somewhere to check filetypes are match
				 * a.ContentType === http.DetectContentType
				 */
				e := Email{
					Filename:    a.Filename,
					Date:        msg.Date(),
					Data:        a.Data,
					Subject:     email.Subject,
					ThreadId:    msg.ThreadID(),
					From:        msg.Header("From"),
					ContentType: a.ContentType,
				}

				emails = append(emails, e)
			}
		}
	}

	return emails, nil
}
