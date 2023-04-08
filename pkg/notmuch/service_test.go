package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type aMock struct{}

func (s *aMock) Query(query string) (*[]Email, error) {
	// TODO update to use testify mocking module
	var emails []Email
	e := Email{
		// pass the query string back to the test so that we can also determine
		// that the query is being provided to the repository
		Filename: query,
	}

	emails = append(emails, e)

	return &emails, nil
}

func TestFetchEmail(t *testing.T) {
	testFilename := "image.jpg"
	mm := Repository(&aMock{})

	r, err := NewService(mm).Query(testFilename)

	assert.Equal(t, nil, err)
	assert.Equal(t, testFilename, (*r)[0].Filename)
}
