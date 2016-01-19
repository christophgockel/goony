package files

import "bufio"

func StreamContent(file File, lines chan string, stop chan bool) {
	scanner := bufio.NewScanner(file)

	for {
		select {
		case <-stop:
			close(lines)
			return
		default:
			if scanner.Scan() {
				lines <- scanner.Text()
			} else {
				stop <- true
			}
		}
	}
}
