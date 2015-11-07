package config

import (
	"errors"
	"strconv"
)

type Options struct {
	Host string
	Port int
	File string
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

	options.Host = arguments[0]
	options.File = arguments[1]

	return options
}

func optionsForThreeArguments(arguments ...string) Options {
	options := newDefaultOptions()

	options.Host = arguments[0]
	options.Port, _ = strconv.Atoi(arguments[1])
	options.File = arguments[2]

	return options
}

func newDefaultOptions() Options {
	return Options{
		Host: "localhost",
		Port: 8080,
	}
}
