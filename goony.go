package main

import (
	"github.com/christophgockel/goony/config"
	"github.com/christophgockel/goony/files"
	"github.com/christophgockel/goony/request"
	"net/http"
	"os"
)

const ROUTINES = 10

func main() {
	options, _ := config.ParseArguments(os.Args[1:]...)

	linesChannel := make(chan string)
	done := make(chan bool)

	file, _ := os.Open(options.File)
	go files.StreamContent(file, linesChannel)

	for i := 0; i < ROUTINES; i++ {
		go func() {
			for line := range linesChannel {
				result := request.Get(line, options.Host, http.DefaultClient)

				files.Print(result, os.Stdout)
			}

			done <- true
		}()
	}

	for i := 0; i < ROUTINES; i++ {
		<-done
	}
}
