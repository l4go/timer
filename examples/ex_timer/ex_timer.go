package main

import (
	"log"
	"time"

	"github.com/l4go/timer"
)

func main() {
	tm := timer.NewTimer()
	defer tm.Stop()
	log.Printf("Start timer\n")
	defer log.Printf("Stop timer\n")

	for i := 0; i < 5; i++ {
		tm.Start(time.Second)
		<-tm.Recv()
		tm.Stop()

		log.Printf("Received a timer event\n")
	}
}
