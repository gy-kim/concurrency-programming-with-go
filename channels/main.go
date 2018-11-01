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

	////////////////////////////////////////////////////////
	/* Buffered channels
	phrase := "There are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
	*/

	/////////////////////////////////////////////////////////
	/* Closing Channels
	phrase := "There are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}

	ch <- "Hello"

	*/

	///////////////////////////////////////////////////////////////
	/* Ranging Over a Channel */
	phrase := "There are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	for msg := range ch {
		fmt.Print(msg + " ")
	}

}
