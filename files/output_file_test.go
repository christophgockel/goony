package files_test

import (
	"github.com/christophgockel/goony/files"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Files - Output File", func() {
	It("returns the file if it exists", func() {
		prepareFilesystemWithAccessibleFile()

		file, err := files.OutputFile("the-file")

		Expect(file).To(Not(BeNil()))
		Expect(err).To(Not(HaveOccurred()))
	})

	It("returns stdout if no filename given", func() {
		file, _ := files.OutputFile("")

		Expect(file).To(Equal(os.Stdout))
	})

	It("returns an error if the file is not writable", func() {
		prepareFilesystemWithUnaccessibleFile()

		_, err := files.OutputFile("unaccessible-file")

		Expect(err).To(HaveOccurred())
	})
})
