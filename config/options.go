package config

type Options struct {
	UsageHelp bool

	NumberOfRoutines int
	Host             string
	File             string
	OutputFilename   string
	RunEndless       bool
}

func newDefaultOptions() Options {
	return Options{
		NumberOfRoutines: 10,
		Host:             "http://localhost",
	}
}
