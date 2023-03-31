package email

import (
	notmuch "github.com/zenhack/go.notmuch"
)

type Service interface {
	FetchEmail(query string) (*notmuch.Messages, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchEmail(query string) (*notmuch.Messages, error) {
	return s.repository.Fetch(query)
}
