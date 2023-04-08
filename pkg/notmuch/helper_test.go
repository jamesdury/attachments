package email

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getMockEmails() []Email {
	var emails []Email
	emails = append(emails, Email{From: "a"})
	emails = append(emails, Email{From: "a"})
	emails = append(emails, Email{From: "a"})
	emails = append(emails, Email{From: "j"})
	emails = append(emails, Email{From: "j"})
	emails = append(emails, Email{From: "a"})
	emails = append(emails, Email{From: "j"})
	emails = append(emails, Email{From: "b"})
	emails = append(emails, Email{From: "j"})
	emails = append(emails, Email{From: "c"})
	emails = append(emails, Email{From: "d"})
	emails = append(emails, Email{From: "e"})
	emails = append(emails, Email{From: "f"})
	emails = append(emails, Email{From: "g"})
	emails = append(emails, Email{From: "j"})

	return emails
}

func TestGetTopContacts(t *testing.T) {
	top := GetTopContacts(getMockEmails())

	assert.NotEqual(t, "j", len(top[0].From))
	assert.NotEqual(t, "a", len(top[1].From))
	assert.Equal(t, 5, len(top))
}

func TestGroupByDate(t *testing.T) {
	const YYYYMMDD = "2006-01-02"
	a, _ := time.Parse(YYYYMMDD, "2023-01-23")
	b, _ := time.Parse(YYYYMMDD, "2023-02-28")
	c, _ := time.Parse(YYYYMMDD, "2023-03-01")

	var emails []Email
	emails = append(emails, Email{Received: a})
	emails = append(emails, Email{Received: a})
	emails = append(emails, Email{Received: b})
	emails = append(emails, Email{Received: b})
	emails = append(emails, Email{Received: c})

	r := GroupByDate(emails)

	assert.Equal(t, 2, len(r[0].Emails)) // a
	assert.Equal(t, 2, len(r[1].Emails)) // b
	assert.Equal(t, 1, len(r[2].Emails)) // c
}
