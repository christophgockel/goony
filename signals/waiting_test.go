package signals_test

import (
	"github.com/christophgockel/goony/signals"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Waiting for a signal", func() {
	It("emits a true value when a signal is read", func() {
		signalChannel := make(chan os.Signal)
		outputChannel := make(chan bool)

		go signals.WaitForSignal(signalChannel, outputChannel)
		signalChannel <- os.Interrupt

		Expect(<-outputChannel).To(BeTrue())
	})
})
