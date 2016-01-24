package format

import (
	"github.com/christophgockel/goony/request"
	"strconv"
	"time"
)

type NoFormatter struct{}

func (formatter NoFormatter) Status(status request.Status) string {
	return status.String()
}
func (formatter NoFormatter) HttpStatus(status int) string {
	return strconv.Itoa(status)
}
func (formatter NoFormatter) Date(date time.Time) string {
	return date.Format("2006-01-02")
}

func (formatter NoFormatter) Time(time time.Time) string {
	return time.Format("15:04:05.000000000")
}
