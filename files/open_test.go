package files_test

import (
	"github.com/christophgockel/goony/files"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Files - Open", func() {
	It("returns the file if it exists", func() {
		prepareFilesystemWithAccessibleFile()

		file, err := files.Open("existing-file")

		Expect(file).To(Not(BeNil()))
		Expect(err).To(Not(HaveOccurred()))
	})

	It("returns an error if the file does not exist", func() {
		prepareFilesystemWithUnexistingFile()

		_, err := files.Open("unknown-file")

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("does not exist"))
	})

	It("returns an error if the file is not accessible", func() {
		prepareFilesystemWithUnaccessibleFile()

		_, err := files.Open("unaccessible-file")

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("permission denied"))
	})
})
