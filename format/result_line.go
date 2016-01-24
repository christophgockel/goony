package format

import (
	"fmt"
	"github.com/christophgockel/goony/request"
)

func ResultLine(result request.Result, formatter Formatter) string {
	return fmt.Sprintf(
		"%s,%s,%s,%s,%s,%d,%s,%s,%s\n",
		formatter.Status(result.Status),
		formatter.Date(result.StartTime),
		formatter.Time(result.StartTime),
		result.Url,
		formatter.HttpStatus(result.HttpStatus),
		result.Nanoseconds,
		formatter.Date(result.EndTime),
		formatter.Time(result.EndTime),
		result.StatusMessage)
}
