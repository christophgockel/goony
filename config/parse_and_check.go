package config

func ParseAndCheck(arguments ...string) (Options, error) {
	options, err := Parse(arguments...)

	if err == nil {
		err = Check(options)
	}

	return options, err
}
