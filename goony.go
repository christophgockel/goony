package main

import (
	"github.com/christophgockel/goony/config"
	"github.com/christophgockel/goony/files"
	"github.com/christophgockel/goony/request"
	"net/http"
	"os"
)

func main() {
	options, _ := config.ParseArguments(os.Args[1:]...)

	file, _ := os.Open(options.File)
	lines := files.ReadContent(file)

	for _, line := range lines {
		result := request.Get(line, options.Host, http.DefaultClient)

		files.Print(result, os.Stdout)
	}
}
