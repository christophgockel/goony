package signals

import "os"

func WaitForSignal(signals chan os.Signal, trues chan bool) {
	<-signals
	trues <- true
}
