package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	testFilename := "image.jpg"
	mm := Repository(&aMock{})

	response, _ := NewService(mm).FetchEmail(testFilename)

	assert.Equal(t, testFilename, response[0].Filename)
}
