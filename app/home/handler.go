package home

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"

	email "github.com/jamesdury/attachments/pkg/notmuch"
)

func GetAttachments(notmuch email.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := strings.Builder{}
		query.WriteString("3months..today")

		q := fmt.Sprintf("attachment:* and date:%s", query.String())

		emails, err := notmuch.Query(q)

		if err != nil {
			return c.Render("static/template/error", fiber.Map{
				"Error": err.Error(),
			})
		}

		return c.Render("static/template/index", fiber.Map{
			"Title":         "Attachments",
			"Emails":        emails,
			"Top":           email.GetTopContacts(*emails),
			"GroupedByDate": email.GroupByDate(*emails),
		})
	}
}

func GetMedia(notmuch email.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		threadid := c.Params("threadid")
		filename := c.Params("filename")

		q := strings.Builder{}
		path, _ := url.PathUnescape(filename)
		// Can just search with "attachment:<filename>" but the search is
		// unreliable via notmuch (if the name contains spaces), so better to
		// use the thread source anyway
		q.WriteString(fmt.Sprintf("thread:%s", threadid))

		thread, err := notmuch.Query(q.String())
		if err != nil {
			return c.Render("static/template/error", fiber.Map{
				"Error": err.Error(),
			})
		}

		var image email.Email
		for _, m := range *thread {
			if m.Filename == path {
				image = m
				break
			}
		}

		c.Status(fiber.StatusOK)
		c.Set(fiber.HeaderContentType, image.ContentType)
		c.SendStream(image.Data)

		return nil
	}
}
