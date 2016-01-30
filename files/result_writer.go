package files

import (
	"github.com/christophgockel/goony/format"
	"github.com/christophgockel/goony/request"
	"io"
)

func Write(result request.Result, output io.Writer, formatter format.Formatter) {
	line := format.ResultLine(result, formatter)

	io.WriteString(output, line)
}
