package email

import (
	"fmt"

	notmuch "github.com/zenhack/go.notmuch"
)

type Repository interface {
	Fetch(query string) (*notmuch.Messages, error)
}
type repository struct {
	connection *notmuch.DB
}

func NewRepo(db *notmuch.DB) Repository {
	return &repository{
		connection: db,
	}
}

func (r *repository) Fetch(query string) (*notmuch.Messages, error) {
	q := r.connection.NewQuery(query)
	q.AddTagExclude("spam")
	q.SetSortScheme(notmuch.SORT_NEWEST_FIRST)
	msgs, err := q.Messages()

	if err != nil {
		return nil, err
	}
	msg := &notmuch.Message{}
	for msgs.Next(&msg) {
		fmt.Println(msg.Date(), msg.Filename())
	}

	return msgs, nil
}
