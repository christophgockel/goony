package files_test

import (
	"github.com/christophgockel/goony/files"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StreamContent", func() {
	var file fakeFile
	var lineChannel chan string
	var stopChannel chan bool

	BeforeEach(func() {
		lineChannel = make(chan string)
		stopChannel = make(chan bool, 1)

		file = fakeFile{}
	})

	It("reads lines of a files into a channel", func() {
		file.AddLine("first line")
		file.AddLine("second line")
		file.AddLine("third line")

		go files.StreamContent(file, lineChannel, stopChannel)

		Expect(<-lineChannel).To(Equal("first line"))
		Expect(<-lineChannel).To(Equal("second line"))
		Expect(<-lineChannel).To(Equal("third line"))
	})

	It("closes the channel when done reading", func() {
		file.AddLine("the only line")

		go files.StreamContent(file, lineChannel, stopChannel)

		DrainRemainingMessages(lineChannel)
		ExpectToBeClosed(lineChannel)
	})

	It("reads file contents until told to stop", func() {
		file.AddLine("first line")
		file.AddLine("second line")

		go files.StreamContent(file, lineChannel, stopChannel)
		stopChannel <- true

		DrainRemainingMessages(lineChannel)
		ExpectToBeClosed(lineChannel)
	})
})

func DrainRemainingMessages(channel chan string) {
	<-channel
}

func ExpectToBeClosed(channel chan string) {
	_, valueCouldBeRead := <-channel

	Expect(valueCouldBeRead).To(BeFalse())
}
