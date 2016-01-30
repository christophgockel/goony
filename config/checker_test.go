package config_test

import (
	"github.com/christophgockel/goony/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config - Checker", func() {
	It("returns an error if no filename is specified", func() {
		options := config.Options{}
		err := config.Check(options)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Filename is missing"))
	})

	It("returns an error if colors are used with an output file", func() {
		options := config.Options{
			File:           "routes-file",
			UseColors:      true,
			OutputFilename: "output-file",
		}

		err := config.Check(options)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Colored output not available in conjunction with an output file"))
	})
})
