package main

import (
	"fmt"
	"math/rand"
	"time"
)

func threading_1() {
	// threading_1 START OMIT
	worker := func(name string) {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(rand.Int63()%1000) * time.Millisecond)
			fmt.Println("hello from", name)
		}
		fmt.Println(name, "done")
	}

	go worker("foo")
	go worker("bar")

	time.Sleep(20 * time.Second)
	// threading_1 END OMIT
}

func sync_1() {
	// sync_1 START OMIT

	// sync_1 END OMIT
}

func main() {
	threading_1()
	sync_1()
}
