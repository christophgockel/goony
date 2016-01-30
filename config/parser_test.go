package config_test

import (
	"github.com/christophgockel/goony/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config - Parser", func() {
	It("returns an error for mistyped flag", func() {
		_, err := config.Parse("-t800")

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Invalid argument: -t800"))
	})

	Context("filename argument", func() {
		It("parses the filename", func() {
			options, _ := config.Parse("filename")

			Expect(options.File).To(Equal("filename"))
		})
	})

	Context("--threads flag", func() {
		It("parses threads (short flag)", func() {
			options, _ := config.Parse("-t", "1000", "filename")

			Expect(options.NumberOfRoutines).To(Equal(1000))
		})

		It("parses threads (long flag)", func() {
			options, _ := config.Parse("--threads", "1000", "filename")

			Expect(options.NumberOfRoutines).To(Equal(1000))
		})

		It("returns an error for invalid thread count", func() {
			_, err := config.Parse("-t", "one", "filename")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Invalid thread count: one"))
		})
	})

	Context("--host flag", func() {
		It("returns an error if filename can't be distinguished", func() {
			_, err := config.Parse("1000", "filename")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Too many arguments"))
		})

		It("parses the host (short flag)", func() {
			options, _ := config.Parse("-h", "http://hostname", "filename")

			Expect(options.Host).To(Equal("http://hostname"))
		})

		It("parses the host (long flag)", func() {
			options, _ := config.Parse("--host", "http://hostname", "filename")

			Expect(options.Host).To(Equal("http://hostname"))
		})

		It("adds HTTP as the default scheme, if not given", func() {
			options, _ := config.Parse("--host", "hostname", "filename")

			Expect(options.Host).To(Equal("http://hostname"))
		})

		It("returns an error for missing hostname", func() {
			_, err := config.Parse("-h")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Missing hostname"))
		})
	})

	Context("--out flag", func() {
		It("parses the filename (short flag)", func() {
			options, _ := config.Parse("-o", "output-filename")

			Expect(options.OutputFilename).To(Equal("output-filename"))
		})

		It("parses the filename (long flag)", func() {
			options, _ := config.Parse("--out", "output-filename")

			Expect(options.OutputFilename).To(Equal("output-filename"))
		})

		It("returns an error for missing filename", func() {
			_, err := config.Parse("--out")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Missing output filename"))
		})
	})

	Context("--endless flag", func() {
		It("parses the short flag", func() {
			options, _ := config.Parse("-e")

			Expect(options.RunEndless).To(BeTrue())
		})

		It("parses the long flag", func() {
			options, _ := config.Parse("--endless")

			Expect(options.RunEndless).To(BeTrue())
		})
	})

	Context("--help flag", func() {
		It("parses the long flag", func() {
			options, _ := config.Parse("--help")

			Expect(options.UsageHelp).To(BeTrue())
		})

		It("ignores any other flags", func() {
			options, err := config.Parse("--endless", "--help", "--threads", "42")

			Expect(err).ToNot(HaveOccurred())
			Expect(options.UsageHelp).To(BeTrue())
			Expect(options.RunEndless).To(BeFalse())
			Expect(options.NumberOfRoutines).ToNot(Equal(42))
		})
	})

	Context("--color flag", func() {
		It("parses short flag", func() {
			options, _ := config.Parse("-c")

			Expect(options.UseColors).To(BeTrue())
		})

		It("parses the long flag", func() {
			options, _ := config.Parse("--color")

			Expect(options.UseColors).To(BeTrue())
		})
	})

	Context("all arguments", func() {
		It("parses all options", func() {
			options, err := config.Parse("-t", "42", "-h", "http://hostname", "-o", "output", "-e", "-c", "filename")

			Expect(err).ToNot(HaveOccurred())
			Expect(options.Host).To(Equal("http://hostname"))
			Expect(options.NumberOfRoutines).To(Equal(42))
			Expect(options.File).To(Equal("filename"))
			Expect(options.OutputFilename).To(Equal("output"))
			Expect(options.RunEndless).To(BeTrue())
			Expect(options.UseColors).To(BeTrue())
		})

		It("doesn't care about the order of the filename and flags", func() {
			options, err := config.Parse("-e", "-c", "filename", "-t", "42", "-h", "http://hostname", "-o", "output")

			Expect(err).ToNot(HaveOccurred())
			Expect(options.Host).To(Equal("http://hostname"))
			Expect(options.NumberOfRoutines).To(Equal(42))
			Expect(options.File).To(Equal("filename"))
			Expect(options.OutputFilename).To(Equal("output"))
			Expect(options.RunEndless).To(BeTrue())
			Expect(options.UseColors).To(BeTrue())
		})
	})
})
