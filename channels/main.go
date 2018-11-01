package main

import "fmt"

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
	/* Ranging Over a Channel

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

	*/

	/////////////////////////////////////////////////////
	/* Switching Between Channels */

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"frodo@underhill,me"},
		From:    "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe.",
	}

	failedMessage := FailedMessage{
		ErrorMessage:    "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedMsg := <-errCh:
		fmt.Println(receivedMsg)
	default:
		fmt.Println("No messages received")
	}

}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
