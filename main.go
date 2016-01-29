package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jpemberthy/jeeves/jeeves"
)

const powerLog = "/Applications/Hearthstone/Logs/Power.log"

func main() {
	ch := make(chan string)
	go jeeves.Watch(powerLog, ch)

	for {
		select {
		case line, ok := <-ch:
			if !ok {
				log.Fatal("error on channel %v", ch)
			}
			if result, err := jeeves.ParseGameResult(line); err == nil {
				fmt.Println(result)
			}
		default:
			time.Sleep(10 * time.Millisecond)
			continue
		}
	}
}
