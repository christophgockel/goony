package config

import (
	"errors"
)

func Check(options Options) error {
	if options.UsageHelp {
		return nil
	}

	if options.File == "" {
		return errors.New("Filename is missing")
	} else if options.UseColors && options.OutputFilename != "" {
		return errors.New("Colored output not available in conjunction with an output file")
	}

	return nil
}
