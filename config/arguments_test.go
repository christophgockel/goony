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
			Expect(options.Host).To(Equal("http://localhost"))
		})
	})

	Context("two arguments", func() {
		var err error

		BeforeEach(func() {
			options, err = config.ParseArguments("http://host", "filename")
		})

		It("takes the host as the first argument", func() {
			Expect(options.Host).To(Equal("http://host"))
		})

		It("uses second argument as the filename", func() {
			Expect(options.File).To(Equal("filename"))
		})

		It("adds 'http://' as the default scheme, if not given", func() {
			options, _ := config.ParseArguments("host.name", "")

			Expect(options.Host).To(Equal("http://host.name"))
		})

		It("adds 'http://' as the default scheme for a hostname with a port", func() {
			options, _ := config.ParseArguments("host:1234", "")

			Expect(options.Host).To(Equal("http://host:1234"))
		})

		It("allows the port to be specified", func() {
			options, _ := config.ParseArguments("http://host.name:1234", "")

			Expect(options.Host).To(Equal("http://host.name:1234"))
		})
	})
})
