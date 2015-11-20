package config

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(arguments ...string) (Options, error) {
	var err error
	options := newDefaultOptions()

	options, err = parseArguments(options, arguments...)

	if err != nil {
		return options, err
	}

	if options.File == "" {
		return options, errors.New("Filename is missing")
	}

	return options, err
}

func parseArguments(options Options, arguments ...string) (Options, error) {
	if len(arguments) == 0 {
		return options, nil
	}

	nextArgument := arguments[0]

	if isFlag(nextArgument) {
		return parseFlag(options, arguments...)
	}

	return parseNonFlagArgument(options, arguments...)
}

func parseFlag(options Options, arguments ...string) (Options, error) {
	flag := arguments[0]

	if isThreadsFlag(flag) {
		return parseThreadArgument(options, arguments...)
	} else if isHostFlag(flag) {
		return parseHostnameArgument(options, arguments...)
	}

	return options, errors.New("Invalid argument: " + flag)
}

func parseNonFlagArgument(options Options, arguments ...string) (Options, error) {
	if fileArgumentIsAllowed(options) {
		options.File = arguments[0]
		return parseArguments(options, arguments[1:]...)
	}

	return options, errors.New("Too many arguments")
}

func isFlag(argument string) bool {
	return strings.HasPrefix(argument, "-")
}

func isThreadsFlag(argument string) bool {
	return argument == "-t" || argument == "--threads"
}

func isHostFlag(argument string) bool {
	return argument == "-h" || argument == "--host"
}

func fileArgumentIsAllowed(options Options) bool {
	return options.File == ""
}

func parseThreadArgument(options Options, arguments ...string) (Options, error) {
	if len(arguments) < 2 {
		return options, errors.New("Missing thread count")
	}

	number, err := strconv.Atoi(arguments[1])

	if err != nil {
		return options, errors.New("Invalid thread count: " + arguments[1])
	}

	options.NumberOfRoutines = number

	return parseArguments(options, arguments[2:]...)
}

func parseHostnameArgument(options Options, arguments ...string) (Options, error) {
	if len(arguments) < 2 {
		return options, errors.New("Missing hostname")
	}

	options.Host = hostWithScheme(arguments[1])

	return parseArguments(options, arguments[2:]...)
}

func hostWithScheme(host string) string {
	if strings.Contains(host, "//") {
		return host
	}

	return "http://" + host
}
