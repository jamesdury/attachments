package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileType(t *testing.T) {
	assert.Equal(t, "pdf", TemplateFunctionFileType("application/pdf"))
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, "Ex", TemplateFunctionTruncate("Example", 2))
}

func TestContact(t *testing.T) {
	assert.Equal(t, "Example", TemplateFunctionContact("Example <example@example.com>"))
	assert.Equal(t, "example@example.com", TemplateFunctionContact("example@example.com"))
}

func TestEmail(t *testing.T) {
	assert.Equal(t, "example@example.com", TemplateFunctionEmail("Example <example@example.com>"))
	assert.Equal(t, "example@example.com", TemplateFunctionEmail("example@example.com"))
	assert.Equal(t, "", TemplateFunctionEmail(""))
}

func TestGravatar(t *testing.T) {
	e := "https://www.gravatar.com/avatar/23463b99b62a72f26ed677cc556c44e8.jpg"

	assert.Equal(t, e, TemplateFunctionGravatar("example@example.com"))

}

func TestPrettyDate(t *testing.T) {
	const YYYYMMDD = "2006-01-02"
	d, _ := time.Parse(YYYYMMDD, "2023-01-04")

	assert.Equal(t, "Jan 04 2023", TemplateFunctionPrettyDate(d))
}
