package format

import (
	"github.com/christophgockel/goony/request"
	"time"
)

type Formatter interface {
	Status(request.Status) string
	HttpStatus(int) string
	Date(time.Time) string
	Time(time.Time) string
}
