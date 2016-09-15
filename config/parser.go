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

	if err != nil || options.UsageHelp {
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
	} else if isOutFlag(flag) {
		return parseOutputFileArgument(options, arguments...)
	} else if isEndlessFlag(flag) {
		return parseEndlessArgument(options, arguments...)
	} else if isHelpFlag(flag) {
		return parseHelpArgument(options, arguments...)
	} else if isCredentialsFlag(flag) {
		return parseCredentialsArgument(options, arguments...)
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

func isOutFlag(argument string) bool {
	return argument == "-o" || argument == "--out"
}

func isEndlessFlag(argument string) bool {
	return argument == "-e" || argument == "--endless"
}

func isHelpFlag(argument string) bool {
	return argument == "--help"
}

func isCredentialsFlag(flag string) bool {
	return flag == "-c" || flag == "--credentials"
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

	options.Host = cleanedHost(arguments[1])

	return parseArguments(options, arguments[2:]...)
}

func cleanedHost(host string) string {
	return hostWithScheme(hostWithoutTrailingSlash(host))
}

func hostWithScheme(host string) string {
	if strings.Contains(host, "//") {
		return host
	}

	return "http://" + host
}

func hostWithoutTrailingSlash(host string) string {
	return strings.TrimRight(host, "/")
}

func parseOutputFileArgument(options Options, arguments ...string) (Options, error) {
	if len(arguments) < 2 {
		return options, errors.New("Missing output filename")
	}

	options.OutputFilename = arguments[1]

	return parseArguments(options, arguments[2:]...)
}

func parseEndlessArgument(options Options, arguments ...string) (Options, error) {
	options.RunEndless = true

	return parseArguments(options, arguments[1:]...)
}

func parseHelpArgument(options Options, arguments ...string) (Options, error) {
	resettedOptions := newDefaultOptions()
	resettedOptions.UsageHelp = true

	return parseArguments(resettedOptions, arguments[0:0]...)
}

func parseCredentialsArgument(options Options, arguments ...string) (Options, error) {
	if len(arguments) == 1 {
		return options, errors.New("Missing username and password")
	}

	credentials := strings.Split(arguments[1], ":")

	if len(credentials) < 2 {
		return options, errors.New("Incomplete credentials provided")
	}

	if credentials[0] == "" {
		return options, errors.New("Incomplete credentials: missing username")
	}

	if len(credentials) == 2 && credentials[1] == "" {
		return options, errors.New("Incomplete credentials: missing password")
	}

	options.Username = credentials[0]
	options.Password = strings.Join(credentials[1:], ":")

	return parseArguments(options, arguments[2:]...)
}
