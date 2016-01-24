package files_test

import (
	"bytes"
	"github.com/christophgockel/goony/files"
	"github.com/christophgockel/goony/format"
	"github.com/christophgockel/goony/request"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Writing of results", func() {
	It("writes results to file", func() {
		outputFile := new(bytes.Buffer)

		result := request.Result{
			Status:        request.SUCCESS,
			StartTime:     time.Date(2015, 11, 13, 15, 06, 01, 123456789, time.Local),
			Url:           "http://host/route/endpoint",
			HttpStatus:    200,
			Nanoseconds:   1522643,
			EndTime:       time.Date(2015, 11, 13, 15, 06, 02, 123456789, time.Local),
			StatusMessage: "",
		}

		files.Write(result, outputFile, format.NoFormatter{})

		Expect(outputFile.String()).To(Equal("S,2015-11-13,15:06:01.123456789,http://host/route/endpoint,200,1522643,2015-11-13,15:06:02.123456789,\n"))
	})
})
