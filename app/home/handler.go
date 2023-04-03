package home

import (
	"sort"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"

	email "github.com/jamesdury/attachments/pkg/notmuch"
)

func getTopContacts(emails []email.Email) []email.Email {
	output := make(map[string][]email.Email)
	// take the keys and put them in a key/value [email.From]: [email..]
	for _, m := range emails {
		output[m.From] = append(output[m.From], m)
	}

	keys := make([]string, 0, len(output))
	for key := range output {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return len(output[keys[i]]) > len(output[keys[j]]) })

	sliced := keys[0:5]
	var newEmails []email.Email

	for _, v := range sliced {
		idx := slices.IndexFunc(emails, func(c email.Email) bool { return c.From == v })
		newEmails = append(newEmails, emails[idx])
	}

	return newEmails
}

func GetAttachments(service email.Service) fiber.Handler {
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
			"Top":    getTopContacts(emails),
		})
	}
}
