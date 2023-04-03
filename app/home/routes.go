package home

import (
	"github.com/gofiber/fiber/v2"
	email "github.com/jamesdury/attachments/pkg/notmuch"
)

func Router(app fiber.Router, service email.Service) {
	app.Get("/", GetAttachments(service))
}
