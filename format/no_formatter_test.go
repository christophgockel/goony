package format_test

import (
	"github.com/christophgockel/goony/format"
	"github.com/christophgockel/goony/request"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("NoFormatter", func() {
	var formatter format.Formatter

	BeforeEach(func() {
		formatter = format.NoFormatter{}
	})

	It("formats successful request status", func() {
		success := request.SUCCESS
		Expect(formatter.Status(success)).To(Equal(success.String()))
	})

	It("formats successful request status", func() {
		failure := request.FAILURE
		Expect(formatter.Status(failure)).To(Equal(failure.String()))
	})

	It("formats HTTP status", func() {
		Expect(formatter.HttpStatus(0)).To(Equal("0"))
		Expect(formatter.HttpStatus(200)).To(Equal("200"))
	})

	It("formats dates", func() {
		date := time.Date(2016, 1, 23, 0, 0, 0, 0, time.Local)

		Expect(formatter.Date(date)).To(Equal("2016-01-23"))
	})

	It("formats times", func() {
		time := time.Date(0, 0, 0, 16, 20, 42, 123456789, time.Local)

		Expect(formatter.Time(time)).To(Equal("16:20:42.123456789"))
	})
})
