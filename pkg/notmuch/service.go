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



func (s *service) Query(query string) (*[]Email, error) {
	return s.repository.Query(query)
}
