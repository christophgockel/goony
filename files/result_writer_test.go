package files_test

import (
	"bytes"
	"github.com/christophgockel/goony/files"
	"github.com/christophgockel/goony/format"
	"github.com/christophgockel/goony/request"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
	"time"
)

var _ = Describe("Result writer", func() {
	var successResult request.Result
	var failureResult request.Result
	var output *bytes.Buffer

	BeforeEach(func() {
		output = new(bytes.Buffer)

		successResult = request.Result{
			Status:        request.SUCCESS,
			StartTime:     time.Date(2015, 11, 13, 15, 06, 01, 123456789, time.Local),
			Url:           "http://host/route/endpoint",
			HttpStatus:    200,
			Nanoseconds:   1522643,
			EndTime:       time.Date(2015, 11, 13, 15, 06, 02, 123456789, time.Local),
			StatusMessage: "",
		}

		failureResult = request.Result{
			Status:        request.FAILURE,
			StartTime:     time.Date(2015, 11, 13, 15, 06, 01, 123456789, time.Local),
			Url:           "http://host/route/endpoint",
			HttpStatus:    0,
			Nanoseconds:   1522643,
			EndTime:       time.Date(2015, 11, 13, 15, 06, 02, 123456789, time.Local),
			StatusMessage: "the failure message",
		}
	})

	It("prints a Result structure as a line of CSV", func() {
		files.Print(successResult, output, format.NoFormatter{})

		Expect(output.String()).To(Equal("S,2015-11-13,15:06:01.123456789,http://host/route/endpoint,200,1522643,2015-11-13,15:06:02.123456789,\n"))
	})

	It("prints a failure Result as a CSV line", func() {
		files.Print(failureResult, output, format.NoFormatter{})

		Expect(output.String()).To(Equal("F,2015-11-13,15:06:01.123456789,http://host/route/endpoint,0,1522643,2015-11-13,15:06:02.123456789,the failure message\n"))
	})

	Context("Formatted Results", func() {
		It("prints results formatted", func() {
			files.Print(successResult, output, DummyFormatter{})

			Expect(output.String()).To(Equal("[S],2015-11-13,15:06:01.123456789,http://host/route/endpoint,<200>,1522643,2015-11-13,15:06:02.123456789,\n"))
		})
	})
})

type DummyFormatter struct {
	format.NoFormatter
}

func (formatter DummyFormatter) Status(status request.Status) string {
	return "[" + status.String() + "]"
}

func (formatter DummyFormatter) HttpStatus(status int) string {
	return "<" + strconv.Itoa(status) + ">"
}
