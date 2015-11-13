package files

import (
	"fmt"
	"github.com/christophgockel/goony/request"
	"io"
)

func Print(result request.Result, output io.Writer) {
	startDate := result.Time.Format("2006-01-02")
	startTime := result.Time.Format("15:04:05.000000000")

	line := fmt.Sprintf("%s,%s,%s,%s,%d,%d,%s\n", result.Status, startDate, startTime, result.Url, result.HttpStatus, result.Nanoseconds, result.StatusMessage)

	io.WriteString(output, line)
}
