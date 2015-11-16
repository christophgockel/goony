package files

import (
	"bufio"
	"io"
)

func StreamContent(input io.Reader, channel chan string) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		channel <- scanner.Text()
	}

	close(channel)
}
