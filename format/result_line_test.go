package format_test

import (
	"github.com/christophgockel/goony/format"
	"github.com/christophgockel/goony/request"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
	"time"
)

var _ = Describe("ResultLine", func() {
	var successResult request.Result
	var failureResult request.Result

	BeforeEach(func() {
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

	It("formats a sucessful result into a line of comma separated values", func() {
		line := format.ResultLine(successResult, format.NoFormatter{})

		Expect(line).To(Equal("S,2015-11-13,15:06:01.123456789,http://host/route/endpoint,200,1522643,2015-11-13,15:06:02.123456789,\n"))
	})

	It("formats a failed result into a line of comma separated values", func() {
		line := format.ResultLine(failureResult, format.NoFormatter{})

		Expect(line).To(Equal("F,2015-11-13,15:06:01.123456789,http://host/route/endpoint,0,1522643,2015-11-13,15:06:02.123456789,the failure message\n"))
	})

	Context("Formatted Results", func() {
		It("prints results formatted", func() {
			line := format.ResultLine(successResult, DummyFormatter{})

			Expect(line).To(Equal("[S],2015-11-13,15:06:01.123456789,http://host/route/endpoint,<200>,1522643,2015-11-13,15:06:02.123456789,\n"))
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
