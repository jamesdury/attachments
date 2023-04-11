package helpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileType(t *testing.T) {
	assert.Equal(t, "pdf", FileType("application/pdf"))
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, "Ex", Truncate("Example", 2))
}

func TestContact(t *testing.T) {
	assert.Equal(t, "Example", Contact("Example <example@example.com>"))
	assert.Equal(t, "example@example.com", Contact("example@example.com"))
}

func TestEmail(t *testing.T) {
	assert.Equal(t, "example@example.com", Email("Example <example@example.com>"))
	assert.Equal(t, "example@example.com", Email("example@example.com"))
	assert.Equal(t, "", Email(""))
}

func TestGravatar(t *testing.T) {
	e := "https://www.gravatar.com/avatar/23463b99b62a72f26ed677cc556c44e8.jpg"

	assert.Equal(t, e, Gravatar("example@example.com"))

}

func TestPrettyDate(t *testing.T) {
	const YYYYMMDD = "2006-01-02"
	d, _ := time.Parse(YYYYMMDD, "2023-01-04")

	assert.Equal(t, "Jan 04 2023", PrettyDate(d))
}

func TestHyponate(t *testing.T) {
	const s = "string with spaces"
	const expected = "string-with-spaces"
	assert.Equal(t, expected, Hyponate(s))
}
