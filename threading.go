package main

import (
	"fmt"
	"math/rand"
	"time"
)

func _1() {
	// _1 START OMIT
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
	// _1 END OMIT
}

func main() {
	_1()
}
