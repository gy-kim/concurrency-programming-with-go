package main

import (
	"fmt"
	"strings"
)

func main() {
	/*  Basic channel
	ch := make(chan string, 1)
	ch <- "Hello"

	fmt.Println(<-ch)
	*/

	phrase := "There are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
}
