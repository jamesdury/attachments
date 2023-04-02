package email

type Service interface {
	FetchEmail(query string) ([]Email, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchEmail(query string) ([]Email, error) {
	return s.repository.Fetch(query)
}
