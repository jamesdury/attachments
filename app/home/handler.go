package home

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

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

type DateEmail struct {
	Date   string
	Emails []email.Email
}

func findDateEmail(output map[int]DateEmail, d string) (DateEmail, int) {
	for i, v := range output {
		if v.Date == d {
			return v, i
		}
	}
	return DateEmail{Date: d}, -1
}

func groupByDate(emails []email.Email) map[int]DateEmail {
	output := make(map[int]DateEmail)

	// take the keys and put them in a key/value [email.From]: [email..]
	for _, m := range emails {
		// slim the date down so that we are specific on the day, rather than
		// day/time
		d := m.Received.Format("Jan 02 2006")
		v, i := findDateEmail(output, d)
		v.Emails = append(v.Emails, m)

		if i == -1 {
			output[len(output)] = v
		} else {
			output[i] = v
		}
	}

	return output
}

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
			"Top":           getTopContacts(*emails),
			"GroupedByDate": groupByDate(*emails),
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
