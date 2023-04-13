package helpers

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

func FileType(s string) string {
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

func Truncate(s string, l int) string {
	return fmt.Sprintf("%.*s", l, s)
}

func Contact(s string) string {
	re := regexp.MustCompile(`(.*?)\ <`)

	m := re.FindStringSubmatch(s)
	if len(m) == 0 {
		return s
	}

	// TODO this needs some checks
	return strings.Trim(m[1], "\"")
}

func Email(s string) string {
	re := regexp.MustCompile(`<(.*)>`)

	m := re.FindStringSubmatch(s)
	if len(m) == 0 {
		return s
	}

	return m[1]
}

func Gravatar(s string) string {
	email := strings.TrimSpace(s)
	email = strings.ToLower(email)

	hash := md5.Sum([]byte(email))
	v := hex.EncodeToString(hash[:])

	return fmt.Sprintf("https://www.gravatar.com/avatar/%s.jpg", v)
}

func PrettyDate(t time.Time) string {
	return t.Format("Jan 02 2006")
}

func Bytesize(b int) string {
	bytes := bytesize.New(float64(b))
	return fmt.Sprintf("%s", bytes)
}

func Escape(s string) string {
	return url.PathEscape(s)
}

func Hyponate(s string) string {
	return strings.Replace(s, " ", "-", -1)
}

func IsImage(s string) bool {
	c := []string{
		"image/gif",
		"image/jpeg",
		"image/png",
		"image/webm",
	}
	for _, v := range c {
		if v == s {
			return true
		}
	}

	return false

}
