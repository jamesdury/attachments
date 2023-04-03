package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/revidian-cloud/go-gravatar/gravatar"
)

func TemplateFunctionFileType(s string) string {
	m := map[string]string{
		"application/pdf": "pdf",
		"application/ics": "ics",
		"image/jpeg":      "jpg",
		"image/png":       "png",
	}
	value, exists := m[s]

	if exists == false {
		return "unsupported"
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
