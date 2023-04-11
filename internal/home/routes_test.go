package home

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	email "github.com/jamesdury/attachments/pkg/notmuch"
)

type aMock struct {
	mock.Mock
}

func (s *aMock) Query(query string) (*[]email.Email, error) {
	var emails []email.Email

	r := strings.NewReader("dummy-image")
	emails = append(emails, email.Email{
		Data:     r,
		Filename: "attachment-name.jpg",
		From:     "a",
	})
	emails = append(emails, email.Email{From: "b", Data: r})
	emails = append(emails, email.Email{From: "c", Data: r})
	emails = append(emails, email.Email{From: "d", Data: r})
	emails = append(emails, email.Email{From: "e", Data: r})
	emails = append(emails, email.Email{From: "f", Data: r})

	return &emails, nil
}

func MockEngineFuncDate(t time.Time) string {
	return ""
}

func MockEngineFuncTruncate(s string, l int) string {
	return ""
}

func MockEngineFuncInt(i int) string {
	return ""
}

func MockEngineFuncString(s string) string {
	return ""
}

func TestHomeRoute(t *testing.T) {

	engine := html.New("../../", ".html") // ../../static/*
	engine.AddFunc("bytesize", MockEngineFuncInt)
	engine.AddFunc("contact", MockEngineFuncString)
	engine.AddFunc("contact", MockEngineFuncString)
	engine.AddFunc("email", MockEngineFuncString)
	engine.AddFunc("escape", MockEngineFuncString)
	engine.AddFunc("filetype", MockEngineFuncString)
	engine.AddFunc("gravatar", MockEngineFuncString)
	engine.AddFunc("hyponate", MockEngineFuncString)
	engine.AddFunc("prettydate", MockEngineFuncDate)
	engine.AddFunc("truncate", MockEngineFuncTruncate)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	mm := email.Repository(&aMock{})

	Router(app, mm)
	req, _ := http.NewRequest(
		"GET",
		"/",
		nil,
	)

	res, err := app.Test(req, -1)

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, res.StatusCode)
}

func TestMediaRoute(t *testing.T) {
	app := fiber.New()

	mm := email.Repository(&aMock{})

	Router(app, mm)
	req, _ := http.NewRequest(
		"GET",
		"/threadid/attachment-name.jpg",
		nil,
	)

	res, err := app.Test(req, -1)

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, res.StatusCode)
	// Read the response body
	body, err := io.ReadAll(res.Body)
	assert.Equal(t, "dummy-image", string(body))
}
