package config_test

import (
	"github.com/christophgockel/goony/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config - ParseAndCheck", func() {
	It("parses and verifies configuration options", func() {
		options, err := config.ParseAndCheck("-h", "http://host.name", "filename")

		Expect(err).ToNot(HaveOccurred())
		Expect(options.Host).To(Equal("http://host.name"))
		Expect(options.File).To(Equal("filename"))
	})

	It("returns errors of argument parsing", func() {
		_, err := config.ParseAndCheck("-h")

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Missing hostname"))
	})

	It("returns errors of option checking", func() {
		_, err := config.ParseAndCheck()

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Filename is missing"))
	})
})
