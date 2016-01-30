package config

import "fmt"

func UsageHelp() string {
	return fmt.Sprintf("Usage: goony [OPTIONS] ROUTES-FILE\n"+
		"\n"+
		"Examples:\n"+
		"    goony routes.txt\n"+
		"    goony --threads 100 routes.txt\n"+
		"    goony --host http://localhost:8080 routes.txt\n"+
		"    goony --endless --host http://example.org routes.txt\n"+
		"\n"+
		"Configuration Options:\n"+
		"    -h, --host          specify the target host (and optional port)\n"+
		"                        (default: %s)\n"+
		"    -t, --threads       define number of parallel threads\n"+
		"                        (default: %d)\n"+
		"    -o, --out FILE      specify target FILE to write results to\n"+
		"    -e, --endless       continuously repeat content of FILE\n"+
		"                        (needs to be stopped with Ctrl+C)\n"+
		"    -c, --color         enable colored output of status codes\n"+
		"        --help          show this usage text", DEFAULT_HOST, DEFAULT_NUMBER_OF_ROUTINES)
}
