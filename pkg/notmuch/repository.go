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
	msgs, err := r.connection.NewQuery(query).Messages()
	if err != nil {
		return nil, err
	}
	fmt.Println(msgs)
	msg := &notmuch.Message{}
	for msgs.Next(&msg) {
		fmt.Println(msg.Filename())
	}

	return msgs, nil
}
