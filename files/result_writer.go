package files

import (
	"fmt"
	"github.com/christophgockel/goony/request"
	"io"
)

func Print(result request.Result, output io.Writer) {
	startDate := result.StartTime.Format("2006-01-02")
	startTime := result.StartTime.Format("15:04:05.000000000")
	endDate := result.EndTime.Format("2006-01-02")
	endTime := result.EndTime.Format("15:04:05.000000000")

	line := fmt.Sprintf("%s,%s,%s,%s,%d,%d,%s,%s,%s\n", result.Status, startDate, startTime, result.Url, result.HttpStatus, result.Nanoseconds, endDate, endTime, result.StatusMessage)

	io.WriteString(output, line)
}
