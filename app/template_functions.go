package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	bytesize "github.com/inhies/go-bytesize"
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

	m := re.FindStringSubmatch(s)
	if len(m) == 0 {
		return s
	}

	// TODO this needs some checks
	return strings.Trim(m[1], "\"")
}

func TemplateFunctionEmail(s string) string {
	re := regexp.MustCompile(`<(.*)>`)

	m := re.FindStringSubmatch(s)
	if len(m) == 0 {
		return s
	}

	return m[1]
}

func TemplateFunctionGravatar(s string) string {
	email := strings.TrimSpace(s)
	email = strings.ToLower(email)

	hash := md5.Sum([]byte(email))
	v := hex.EncodeToString(hash[:])

	return fmt.Sprintf("https://www.gravatar.com/avatar/%s.jpg", v)
}

func TemplateFunctionPrettyDate(t time.Time) string {
	return t.Format("Jan 02 2006")
}

func TemplateFunctionBytesize(b int) string {
	bytes := bytesize.New(float64(b))
	return fmt.Sprintf("%s", bytes)
}

func TemplateFunctionEscape(s string) string {
	return url.PathEscape(s)
}
