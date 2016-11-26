package main

import (
	"log"
	"sync"
	"time"
)

func sync_2() {
	// sync_2 START OMIT

	worker := func(wg *sync.WaitGroup) {
		log.Println("working")
		time.Sleep(5 * time.Second)
		log.Println("almost done")
		wg.Done()
	}

	wg := &sync.WaitGroup{}

	go worker(wg)
	go worker(wg)
	wg.Add(2)

	wg.Wait()
	log.Println("done")
	// sync_2 END OMIT
}

func main() {
	sync_2()
}
