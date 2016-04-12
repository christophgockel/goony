package config

const DEFAULT_NUMBER_OF_ROUTINES = 10
const DEFAULT_HOST = "http://localhost"

type Options struct {
	UsageHelp bool

	NumberOfRoutines int
	Host             string
	File             string
	OutputFilename   string
	RunEndless       bool
	Username         string
	Password         string
}

func newDefaultOptions() Options {
	return Options{
		NumberOfRoutines: DEFAULT_NUMBER_OF_ROUTINES,
		Host:             DEFAULT_HOST,
	}
}
