package files_test

import (
	"bytes"
	"github.com/christophgockel/goony/files"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Files", func() {
	It("reads lines of a file into a list of strings", func() {
		file := new(bytes.Buffer)

		file.WriteString("first line\n")
		file.WriteString("second line\n")
		file.WriteString("third line")

		lines := files.ReadContent(file)

		Expect(len(lines)).To(Equal(3))
		Expect(lines[0]).To(Equal("first line"))
		Expect(lines[1]).To(Equal("second line"))
		Expect(lines[2]).To(Equal("third line"))
	})
})
