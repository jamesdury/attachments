package email

import (
	"io"
	"time"
)

type Email struct {
	Data       io.Reader
	Date       time.Time
	Filename   string
	From       string
	Subject    string
	Suspicious bool
	ThreadId   string
	To         string
}
