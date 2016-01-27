package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jpemberthy/jeeves/watcher"
)

const powerLog = "/Applications/Hearthstone/Logs/Power.log"

func main() {
	ch := make(chan string)
	go watcher.Watch(powerLog, ch)

	for {
		select {
		case line, ok := <-ch:
			if !ok {
				log.Fatal("error on channel %v", ch)
			}
			fmt.Println(line)
		default:
			time.Sleep(10 * time.Millisecond)
			continue
		}
	}
}
