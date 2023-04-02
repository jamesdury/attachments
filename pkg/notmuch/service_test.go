package email

import (
	"testing"
)

type aMock struct{}

func (s *aMock) Fetch(query string) ([]Email, error) {
	var emails []Email
	e := Email{
		// pass the query string back to the test so that we can also determine
		// that the query is being provided to the repository
		Filename: query,
	}

	emails = append(emails, e)

	return emails, nil
}

func TestFetchEmail(t *testing.T) {

	mm := Repository(&aMock{})

	var testFilename = "image.jpg"
	response, _ := NewService(mm).FetchEmail(testFilename)

	if response[0].Filename != testFilename {
		t.Fatalf("Service is not calling repository.Fetch()")
	}
}
