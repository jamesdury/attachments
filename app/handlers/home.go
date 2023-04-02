package handlers

import (
	"github.com/gofiber/fiber/v2"

	email "github.com/jamesdury/attachments/pkg/notmuch"
)

func GetEmails(service email.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		emails, err := service.FetchEmail("attachment:* and date:3months..today")

		if err != nil {
			return c.Render("static/template/error", fiber.Map{
				"Error": err.Error(),
			})

		}

		return c.Render("static/template/index", fiber.Map{
			"Title":  "Attachments",
			"Emails": emails,
		})
	}
}
