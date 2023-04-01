package handlers

import (
	"github.com/gofiber/fiber/v2"

	email "com.jamesdury.emailfiles/pkg/notmuch"
)

func GetEmails(service email.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service.FetchEmail("attachment:* and date:3months..today")

		return c.Render("template/index", fiber.Map{
			"Title": "Hello, emails (with attachments)!",
		})
	}
}
