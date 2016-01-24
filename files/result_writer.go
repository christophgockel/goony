package files

import (
	"fmt"
	"github.com/christophgockel/goony/format"
	"github.com/christophgockel/goony/request"
	"io"
)

func Print(result request.Result, output io.Writer, formatter format.Formatter) {
	line := fmt.Sprintf(
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

	io.WriteString(output, line)
}
