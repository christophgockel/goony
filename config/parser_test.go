package config_test

import (
	"github.com/christophgockel/goony/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("Config - Parser", func() {
	It("returns an error if no arguments are given", func() {
		_, err := config.Parse()

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Filename is missing"))
	})

	It("returns an error for mistyped flag", func() {
		_, err := config.Parse("-t800")

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Invalid argument: -t800"))
	})

	Context("filename argument", func() {
		It("returns an error if no file has been specified", func() {
			_, err := config.Parse("-h", "host.name")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Filename is missing"))
		})

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

		It("strips trailing slashes if given", func() {
			options, _ := config.Parse("--host", "http://hostname/", "filename")

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

	Context("--credentials flag", func() {
		It("parses the credentials (short flag)", func() {
			options, _ := config.Parse("-c", "username:password")

			Expect(options.Username).To(Equal("username"))
			Expect(options.Password).To(Equal("password"))
		})

		It("parses the credentials (long flag)", func() {
			options, _ := config.Parse("--credentials", "username:password")

			Expect(options.Username).To(Equal("username"))
			Expect(options.Password).To(Equal("password"))
		})

		It("ignores surplus colons", func() {
			options, _ := config.Parse("-c", "username:password:with:colons")

			Expect(options.Username).To(Equal("username"))
			Expect(options.Password).To(Equal("password:with:colons"))
		})

		It("returns an error if credentials are missing", func() {
			_, err := config.Parse("-c")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Missing username and password"))
		})

		It("returns an error when incomplete credentials are provided", func() {
			_, err := config.Parse("-c", "username")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Incomplete credentials provided"))
		})

		It("returns an error when password is missing", func() {
			_, err := config.Parse("-c", "username:")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Incomplete credentials: missing password"))
		})

		It("returns an error when username is missing", func() {
			_, err := config.Parse("-c", ":password")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Incomplete credentials: missing username"))
		})

		It("allows passwords to begin with a colon", func() {
			options, _ := config.Parse("-c", "username::password")

			Expect(options.Username).To(Equal("username"))
			Expect(options.Password).To(Equal(":password"))
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

	Context("all arguments", func() {
		It("parses all options in any order", func() {
			allPossibleArguments := CommandLineArguments{
				[]string{"-t", "42"},
				[]string{"-h", "http://hostname"},
				[]string{"-o", "output"},
				[]string{"filename"},
				[]string{"-e"},
				[]string{"-c", "username:password"},
			}
			arguments := flatten(shuffle(allPossibleArguments))

			options, err := config.Parse(arguments...)

			Expect(err).ToNot(HaveOccurred())
			Expect(options.Host).To(Equal("http://hostname"))
			Expect(options.NumberOfRoutines).To(Equal(42))
			Expect(options.File).To(Equal("filename"))
			Expect(options.OutputFilename).To(Equal("output"))
			Expect(options.RunEndless).To(BeTrue())
			Expect(options.Username).To(Equal("username"))
			Expect(options.Password).To(Equal("password"))
		})
	})
})

type CommandLineArguments [][]string

func shuffle(arguments CommandLineArguments) CommandLineArguments {
	rand.Seed(time.Now().UnixNano())

	for i := range arguments {
		j := rand.Intn(i + 1)
		arguments[i], arguments[j] = arguments[j], arguments[i]
	}

	return arguments
}

func flatten(arguments CommandLineArguments) []string {
	var flattened []string

	for i := range arguments {
		for j := range arguments[i] {
			flattened = append(flattened, arguments[i][j])
		}
	}

	return flattened
}
