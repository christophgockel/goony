package format_test

import (
	"github.com/christophgockel/goony/format"
	"github.com/christophgockel/goony/request"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ColoredFormatter", func() {
	var formatter format.Formatter

	BeforeEach(func() {
		formatter = format.ColoredFormatter{}
	})

	It("colorises success status", func() {
		success := request.SUCCESS
		Expect(formatter.Status(success)).To(Equal(format.Success(success.String())))
	})

	It("colorises failure status", func() {
		failure := request.FAILURE
		Expect(formatter.Status(failure)).To(Equal(format.Error(failure.String())))
	})

	It("colorises successful HTTP statuses", func() {
		Expect(formatter.HttpStatus(100)).To(Equal(format.Success("100")))
		Expect(formatter.HttpStatus(200)).To(Equal(format.Success("200")))
		Expect(formatter.HttpStatus(308)).To(Equal(format.Success("308")))
	})

	It("colorises unsuccessful HTTP statuses", func() {
		Expect(formatter.HttpStatus(400)).To(Equal(format.Error("400")))
		Expect(formatter.HttpStatus(500)).To(Equal(format.Error("500")))
		Expect(formatter.HttpStatus(0)).To(Equal(format.Error("0")))
	})
})
