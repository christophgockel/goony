package config

import (
	"errors"
	"strconv"
	"strings"
)

type Options struct {
	NumberOfRoutines int
	Host             string
	File             string
}

func ParseArguments(arguments ...string) (Options, error) {
	options := newDefaultOptions()

	if len(arguments) == 1 {
		return optionsForOneArgument(arguments...), nil
	} else if len(arguments) == 2 {
		return optionsForTwoArguments(arguments...), nil
	} else if len(arguments) == 3 {
		return optionsForThreeArguments(arguments...), nil
	}

	return options, errors.New("No filename specified")
}

func optionsForOneArgument(arguments ...string) Options {
	options := newDefaultOptions()

	options.File = arguments[0]

	return options
}

func optionsForTwoArguments(arguments ...string) Options {
	options := newDefaultOptions()

	options.Host = urlWithScheme(arguments[0])
	options.File = arguments[1]

	return options
}

func optionsForThreeArguments(arguments ...string) Options {
	options := newDefaultOptions()

	options.NumberOfRoutines, _ = strconv.Atoi(arguments[0])
	options.Host = urlWithScheme(arguments[1])
	options.File = arguments[2]

	return options
}

func urlWithScheme(originalUrl string) string {
	newUrl := originalUrl

	if strings.Contains(originalUrl, "//") == false {
		newUrl = "http://" + originalUrl
	}

	return newUrl
}

func newDefaultOptions() Options {
	return Options{
		NumberOfRoutines: 10,
		Host:             "http://localhost",
	}
}
