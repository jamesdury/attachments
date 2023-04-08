package email

import (
	"sort"

	"golang.org/x/exp/slices"
)

type Service interface {
	Query(query string) (*[]Email, error)
	GroupByDate(emails []Email) map[int]DateEmail
	GetTopContacts(emails []Email) []Email
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func findDateEmail(output map[int]DateEmail, d string) (DateEmail, int) {
	for i, v := range output {
		if v.Date == d {
			return v, i
		}
	}
	return DateEmail{Date: d}, -1
}

func (s *service) GroupByDate(emails []Email) map[int]DateEmail {
	output := make(map[int]DateEmail)

	// take the keys and put them in a key/value [email.From]: [email..]
	for _, m := range emails {
		// slim the date down so that we are specific on the day, rather than
		// day/time
		d := m.Received.Format("Jan 02 2006")
		v, i := findDateEmail(output, d)
		v.Emails = append(v.Emails, m)

		if i == -1 {
			output[len(output)] = v
		} else {
			output[i] = v
		}
	}

	return output
}

// get the top 5 contacts in an email array
func (s *service) GetTopContacts(emails []Email) []Email {
	output := make(map[string][]Email)
	// take the keys and put them in a key/value [email.From]: [email..]
	for _, m := range emails {
		output[m.From] = append(output[m.From], m)
	}

	keys := make([]string, 0, len(output))
	for key := range output {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return len(output[keys[i]]) > len(output[keys[j]]) })

	sliced := keys[0:5]
	var newEmails []Email

	for _, v := range sliced {
		idx := slices.IndexFunc(emails, func(c Email) bool { return c.From == v })
		newEmails = append(newEmails, emails[idx])
	}

	return newEmails
}


func (s *service) Query(query string) (*[]Email, error) {
	return s.repository.Query(query)
}
