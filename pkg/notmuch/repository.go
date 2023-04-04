package email

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"

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

// https://stackoverflow.com/a/56141678
func Filesize(s string) int {
	l := len(s)

	// count how many trailing '=' there are (if any)
	eq := 0
	if l >= 2 {
		if s[l-1] == '=' {
			eq++
		}
		if s[l-2] == '=' {
			eq++
		}

		l -= eq
	}

	// basically:
	//
	// eq == 0 :    bits-wasted = 0
	// eq == 1 :    bits-wasted = 2
	// eq == 2 :    bits-wasted = 4

	// each base64 character = 6 bits

	// so orig length ==  (l*6 - eq*2) / 8

	return (l*3 - eq) / 4

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

				bytes, err := ioutil.ReadAll(a.Data)
				if err != nil {
					panic(err)
				}

				// TODO Determine the content type of the image file
				// mimeType := http.DetectContentType(bytes)
				b64 := base64.StdEncoding.EncodeToString(bytes)

				/*
				 * TODO write a check somewhere to check filetypes are match
				 * a.ContentType === http.DetectContentType
				 */
				e := Email{
					ContentType: a.ContentType,
					Data:        a.Data,
					Filename:    a.Filename,
					From:        msg.Header("From"),
					Received:    msg.Date(),
					Size:        Filesize(b64),
					Subject:     email.Subject,
					ThreadId:    msg.ThreadID(),
				}

				emails = append(emails, e)
			}
		}
	}

	return &emails, nil
}
