package files_test

import (
	"bytes"
	"github.com/christophgockel/goony/files"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Files", func() {
	It("reads lines of a files into a channel", func() {
		file := new(bytes.Buffer)

		file.WriteString("first line\n")
		file.WriteString("second line\n")
		file.WriteString("third line")

		channel := make(chan string)

		go files.StreamContent(file, channel)

		Expect(<-channel).To(Equal("first line"))
		Expect(<-channel).To(Equal("second line"))
		Expect(<-channel).To(Equal("third line"))
	})

	It("closes the channel when done reading", func() {
		file := new(bytes.Buffer)
		channel := make(chan string)

		file.WriteString("the only line")

		go files.StreamContent(file, channel)

		<-channel

		lastResult, channelIsOpen := <-channel

		Expect(lastResult).To(Equal(""))
		Expect(channelIsOpen).To(Equal(false))
	})
})
