package request_test

import (
	"github.com/christophgockel/goony/request"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request Data", func() {
	It("returns the full request URL", func() {
		data := request.Data{"http://host", "/endpoint", "", ""}

		Expect(data.Url()).To(Equal("http://host/endpoint"))
	})
})
