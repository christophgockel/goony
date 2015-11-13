package files_test

import (
	"bytes"
	"github.com/christophgockel/goony/files"
	"github.com/christophgockel/goony/request"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Result writer", func() {
	It("prints a Result structure as a line of CSV", func() {
		result := request.Result{
			Status:        request.SUCCESS,
			Time:          time.Date(2015, 11, 13, 15, 06, 01, 123456789, time.Local),
			Url:           "http://host/route/endpoint",
			HttpStatus:    200,
			Nanoseconds:   1522643,
			StatusMessage: "",
		}

		output := new(bytes.Buffer)

		files.Print(result, output)

		Expect(output.String()).To(Equal("S,2015-11-13,15:06:01.123456789,http://host/route/endpoint,200,1522643,\n"))
	})
})
