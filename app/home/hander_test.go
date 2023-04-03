package home

import (
	"testing"

	email "github.com/jamesdury/attachments/pkg/notmuch"
)

func TestGetTopContacts(t *testing.T) {
	var emails []email.Email

	emails = append(emails, email.Email{From: "a"})
	emails = append(emails, email.Email{From: "a"})
	emails = append(emails, email.Email{From: "a"})
	emails = append(emails, email.Email{From: "j"})
	emails = append(emails, email.Email{From: "j"})
	emails = append(emails, email.Email{From: "a"})
	emails = append(emails, email.Email{From: "j"})
	emails = append(emails, email.Email{From: "b"})
	emails = append(emails, email.Email{From: "j"})
	emails = append(emails, email.Email{From: "c"})
	emails = append(emails, email.Email{From: "d"})
	emails = append(emails, email.Email{From: "e"})
	emails = append(emails, email.Email{From: "f"})
	emails = append(emails, email.Email{From: "g"})
	emails = append(emails, email.Email{From: "j"})

	top := getTopContacts(emails)

	if top[0].From != "j" {
		t.Fatalf("Emails are not being sorted")
	}
	if top[1].From != "a" {
		t.Fatalf("Emails are not being sorted")
	}
}
