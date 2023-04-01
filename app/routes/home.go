package routes

import (
	email "com.jamesdury.emailfiles/pkg/notmuch"
	"github.com/gofiber/fiber/v2"

	handlers "com.jamesdury.emailfiles/app/handlers"
)

// BookRouter is the Router for GoFiber App
func HomeRouter(app fiber.Router, service email.Service) {
	app.Get("/", handlers.GetEmails(service))
}
