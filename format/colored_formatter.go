package format

import (
	"github.com/christophgockel/goony/request"
	"strconv"
)

type ColoredFormatter struct {
	NoFormatter
}

func (formatter ColoredFormatter) Status(status request.Status) string {
	if status == request.SUCCESS {
		return Success(status.String())
	} else {
		return Error(status.String())
	}
}

func (formatter ColoredFormatter) HttpStatus(status int) string {
	if status == 0 || status >= 400 {
		return Error(strconv.Itoa(status))
	}
	return Success(strconv.Itoa(status))
}

func Success(value string) string {
	return "\x1b[32m" + value + "\x1b[0m"
}

func Error(value string) string {
	return "\x1b[31m" + value + "\x1b[0m"
}
