package main

import (
	"fmt"
	"github.com/christophgockel/goony/config"
	"github.com/christophgockel/goony/files"
	"github.com/christophgockel/goony/request"
	"net/http"
	"os"
)

func main() {
	options, err := config.Parse(os.Args[1:]...)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	linesChannel := make(chan string)
	done := make(chan bool)
	resultsChannel := make(chan request.Result, 10000)

	file, _ := os.Open(options.File)
	go files.StreamContent(file, linesChannel)

	for i := 0; i < options.NumberOfRoutines; i++ {
		go func() {
			for line := range linesChannel {
				resultsChannel <- request.Get(line, options.Host, http.DefaultClient)
			}

			done <- true
		}()
	}

	go func() {
		for result := range resultsChannel {
			files.Print(result, os.Stdout)
		}
		done <- true
	}()

	for i := 0; i < options.NumberOfRoutines; i++ {
		<-done
	}

	close(resultsChannel)
	<-done
}
