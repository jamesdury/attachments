package email

import (
	"io"
	"time"
)

type Email struct {
	ContentType string
	Data        io.Reader
	Filename    string
	Filesize    int
	From        string
	Received    time.Time
	Size        int
	Subject     string
	Suspicious  bool
	ThreadId    string
	To          string
}
