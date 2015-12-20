package files_test

import (
	"github.com/christophgockel/goony/files"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Files - Reading and Writing", func() {
	Context("OpenForReading()", func() {
		It("returns the file if it exists", func() {
			prepareFilesystemWithAccessibleFile()

			file, err := files.OpenForReading("existing-file")

			Expect(file).To(Not(BeNil()))
			Expect(err).To(Not(HaveOccurred()))
		})

		It("returns an error if the file does not exist", func() {
			prepareFilesystemWithUnexistingFile()

			_, err := files.OpenForReading("unknown-file")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("does not exist"))
		})

		It("returns an error if the file is not accessible", func() {
			prepareFilesystemWithUnaccessibleFile()

			_, err := files.OpenForReading("unaccessible-file")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("permission denied"))
		})
	})

	Context("OpenForWriting()", func() {
		It("returns the file if it exists", func() {
			prepareFilesystemWithAccessibleFile()

			file, err := files.OpenForWriting("the-file")

			Expect(file).To(Not(BeNil()))
			Expect(err).To(Not(HaveOccurred()))
		})

		It("returns stdout if no filename given", func() {
			file, _ := files.OpenForWriting("")

			Expect(file).To(Equal(os.Stdout))
		})

		It("returns an error if the file is not writable", func() {
			prepareFilesystemWithUnaccessibleFile()

			_, err := files.OpenForWriting("unaccessible-file")

			Expect(err).To(HaveOccurred())
		})
	})
})
