package files_test

import (
	"bufio"
	"github.com/christophgockel/goony/files"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Endless File", func() {
	var scanner *bufio.Scanner

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	It("continously reads file contents", func() {
		file := fakeFile{}
		file.AddLine("line 1")
		file.AddLine("line 2")

		endlessFile := files.EndlessFile{file}
		scanner = bufio.NewScanner(endlessFile)

		Expect(readLine()).To(Equal("line 1"))
		Expect(readLine()).To(Equal("line 2"))
		Expect(readLine()).To(Equal("line 1"))
		Expect(readLine()).To(Equal("line 2"))
	})
})
