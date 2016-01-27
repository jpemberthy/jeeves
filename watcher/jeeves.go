package watcher

import (
	"bufio"
	"io"
	"os"
	"time"
)

// Watch watches the file present on path and sends each line
// to channel
func Watch(path string, ch chan string) error {
	defer close(ch)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(0, 2) // go to the end of file.
	reader := bufio.NewReader(file)

	ticker := time.NewTicker(500 * time.Millisecond)
	for range ticker.C {
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			ch <- line
		}
	}

	return nil
}
