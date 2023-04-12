// Package email is an interface to communicate with the notmuch server
package email

type Service interface {
	Query(query string) (*[]Email, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Query returns a reference to an array of emails based on the notmuch query
// provided
func (s *service) Query(query string) (*[]Email, error) {
	return s.repository.Query(query)
}
