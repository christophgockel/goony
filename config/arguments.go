package config

import (
	"errors"
	"strings"
)

type Options struct {
	Host string
	File string
}

func ParseArguments(arguments ...string) (Options, error) {
	options := newDefaultOptions()

	if len(arguments) == 1 {
		return optionsForOneArgument(arguments...), nil
	} else if len(arguments) == 2 {
		return optionsForTwoArguments(arguments...), nil
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

func urlWithScheme(originalUrl string) string {
	newUrl := originalUrl

	if strings.Contains(originalUrl, "//") == false {
		newUrl = "http://" + originalUrl
	}

	return newUrl
}

func newDefaultOptions() Options {
	return Options{
		Host: "http://localhost",
	}
}
