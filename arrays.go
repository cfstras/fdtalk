package main

import (
	"fmt"
	"log"
)

func arrays_1() {
	// arrays_1 START OMIT
	array1 := []string{}
	array2 := make([]string, 0)
	var array3 []string

	log.Println("array1:", array1)
	log.Println("array2:", array2)
	log.Println("array3:", array3)

	for i := 0; i < 10; i++ {
		array1 = append(array1, fmt.Sprint(i))
		array2 = append(array2, fmt.Sprint(i))
		array3 = append(array3, fmt.Sprint(i))
	}

	log.Println("array1:", array1)
	log.Println("array2:", array2)
	log.Println("array3:", array3)

	// arrays_1 END OMIT
}

func main() {
	arrays_1()
}
