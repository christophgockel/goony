package files

import (
	"bufio"
	"io"
)

func ReadContent(input io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
