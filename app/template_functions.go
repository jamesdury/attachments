package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/revidian-cloud/go-gravatar/gravatar"
)

func TemplateFunctionFileType(s string) string {
	m := map[string]string{
		"application/ics":              "ics",
		"application/msword":           "docx",
		"application/pdf":              "pdf",
		"application/x-zip-compressed": "zip",
		"image/jpeg":                   "jpg",
		"image/png":                    "png",
		"text/calendar":                "ics",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":       "xlsx",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": "docx",
	}
	value, exists := m[s]

	if exists == false {
		return "blank"
	}

	return value
}

func TemplateFunctionTruncate(s string, l int) string {
	return fmt.Sprintf("%.*s", l, s)
}

func TemplateFunctionContact(s string) string {
	re := regexp.MustCompile(`(.*?)\ <`)
	// TODO this needs some checks
	return strings.Trim(re.FindStringSubmatch(s)[1], "\"")
}

func TemplateFunctionEmail(s string) string {
	re := regexp.MustCompile(`<(.*)>`)
	return re.FindStringSubmatch(s)[1]
}

func TemplateFunctionGravatar(s string) string {
	img, err := gravatar.NewImage(s)
	if err != nil {
		panic(err)
	}

	imgURL, err := img.URL()
	if err != nil {
		panic(err)
	}

	return imgURL.String()
}

func TemplateFunctionPrettyDate(t time.Time) string {
	return t.Format("Jan 02")
}
