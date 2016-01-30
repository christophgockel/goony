package main

import (
	"fmt"
	"github.com/christophgockel/goony/config"
	"github.com/christophgockel/goony/files"
	"github.com/christophgockel/goony/format"
	"github.com/christophgockel/goony/request"
	"github.com/christophgockel/goony/signals"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	options := options()
	routesFile := routesFile(options.File, options.RunEndless)
	outputFile := outputFile(options.OutputFilename)

	defer routesFile.Close()
	defer outputFile.Close()

	linesChannel := make(chan string)
	done := make(chan bool)
	resultsChannel := make(chan request.Result, 10000)

	startContentStream(options, routesFile, linesChannel)

	for i := 0; i < options.NumberOfRoutines; i++ {
		go func() {
			for line := range linesChannel {
				resultsChannel <- request.Get(line, options.Host, http.DefaultClient)
			}

			done <- true
		}()
	}

	go func() {
		formatter := formatter(options.UseColors)

		for result := range resultsChannel {
			files.Write(result, outputFile, formatter)
		}
		done <- true
	}()

	for i := 0; i < options.NumberOfRoutines; i++ {
		<-done
	}

	close(resultsChannel)
	<-done
}

func options() config.Options {
	options, err := config.ParseAndCheck(os.Args[1:]...)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		usageHelp()
		os.Exit(1)
	}

	if options.UsageHelp {
		usageHelp()
		os.Exit(99)
	}

	return options
}

func usageHelp() {
	fmt.Fprintln(os.Stdout, config.UsageHelp())
}

func routesFile(filename string, endless bool) files.File {
	file, err := files.OpenForReading(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	if endless {
		return files.EndlessFile{file}
	} else {
		return file
	}
}

func outputFile(filename string) files.File {
	file, err := files.OpenForWriting(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

	return file
}

func startContentStream(options config.Options, file files.File, linesChannel chan string) {
	stopChannel := make(chan bool, 1)
	catchCtrlC(stopChannel)
	go files.StreamContent(file, linesChannel, stopChannel)
}

func catchCtrlC(output chan bool) {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)

	go signals.WaitForSignal(signalChannel, output)
}

func formatter(useColors bool) format.Formatter {
	if useColors {
		return format.ColoredFormatter{}
	} else {
		return format.NoFormatter{}
	}
}
