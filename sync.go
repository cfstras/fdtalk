package main

import (
	"log"
	"time"
)

func sync_1() {
	// sync_1 START OMIT
	worker := func(quit chan<- bool) {
		log.Println("working")
		time.Sleep(5 * time.Second)
		log.Println("almost done")
		quit <- true
	}

	quit := make(chan bool)
	go worker(quit)

	<-quit
	log.Println("done")
	// sync_1 END OMIT
}

func main() {
	sync_1()
}
