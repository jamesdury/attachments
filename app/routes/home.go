package routes

import (
	"github.com/gofiber/fiber/v2"
	email "github.com/jamesdury/attachments/pkg/notmuch"

	handlers "github.com/jamesdury/attachments/app/handlers"
)

func HomeRouter(app fiber.Router, service email.Service) {
	app.Get("/", handlers.GetEmails(service))
}
