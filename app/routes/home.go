package routes

import (
	email "com.jamesdury.emailfiles/pkg/notmuch"
	"github.com/gofiber/fiber/v2"

	handlers "com.jamesdury.emailfiles/app/handlers"
)

func HomeRouter(app fiber.Router, service email.Service) {
	app.Get("/", handlers.GetEmails(service))
}
