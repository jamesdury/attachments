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


func TestFileSize(t *testing.T) {


	base64 := `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
YII=`
	assert.Equal(t, 2716, filesize(base64))

}
