package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

const watchedLog = "/Applications/Hearthstone/Logs/Power.log"

func main() {
	file, err := os.Open(watchedLog)
	if err != nil {
		log.Fatal(err)
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
				log.Fatal(err)
			}
			log.Println(line)
		}
	}
}
