package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
	}

	//create a channel -> a pipe that allows us to send and receive messages between goroutines
	ch := make(chan string)

	for _, link := range links {

		//launch a goroutine -> a lightweight thread managed by the Go runtime and
		// channel is passed as an argument to the function - this allows the function to send a message back to the main goroutine
		go checkLink(link, ch) // this is a blocking call - it will wait for the function to finish and then send a message back to the main goroutine
		// if we don't pass the channel, the function will not be able to send a message back to the main goroutine
		// the channel is used to communicate between the main and the goroutine

	}
	//receive a message from the channel
	// <- is the receive operator
	msg := <-ch // this is the message from the channel

	// print the message
	fmt.Println(msg)

}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		// fmt.Println("❌", link, "might be down!")
		ch <- "❌ " + link + " might be down!"
		return
	}

	// fmt.Println("✅", link, "is up!")

	// send a message to the channel
	// <- is the send operator
	ch <- "✅ " + link + " is up!" // this is the message we want to send to the channel

	//diff between send and receive is that send is used to send a message to the channel and receive is used to receive a message from the channel
	// but the <- operator can be used to send and receive messages - it depends on the direction of the arrow
	// ch <- "✅" is sending a message to the channel
	// msg := <-ch is receiving a message from the channel

}
