package config_test

import (
	"github.com/christophgockel/goony/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Arguments", func() {
	var options config.Options

	Context("no argument", func() {
		It("returns an error", func() {
			_, err := config.ParseArguments()

			Expect(err).To(HaveOccurred())
			Expect(strings.ToLower(err.Error())).To(ContainSubstring("no filename"))
		})
	})

	Context("one argument", func() {
		BeforeEach(func() {
			options, _ = config.ParseArguments("filename")
		})

		It("is the log file's name", func() {
			Expect(options.File).To(Equal("filename"))
		})

		It("uses localhost as the default host", func() {
			Expect(options.Host).To(Equal("localhost"))
		})

		It("uses 8080 as the default port", func() {
			Expect(options.Port).To(Equal(8080))
		})
	})

	Context("two arguments", func() {
		BeforeEach(func() {
			options, _ = config.ParseArguments("host", "filename")
		})

		It("takes the host as the first argument", func() {
			Expect(options.Host).To(Equal("host"))
		})

		It("uses second argument as the filename", func() {
			Expect(options.File).To(Equal("filename"))
		})

		It("uses 8080 as the default port", func() {
			Expect(options.Port).To(Equal(8080))
		})
	})

	Context("three arguments", func() {
		BeforeEach(func() {
			options, _ = config.ParseArguments("host", "1234", "filename")
		})

		It("takes the host as the first argument", func() {
			Expect(options.Host).To(Equal("host"))
		})

		It("takes the port as the second argument", func() {
			Expect(options.Port).To(Equal(1234))
		})

		It("uses the third argument as the filename", func() {
			Expect(options.File).To(Equal("filename"))
		})
	})
})
