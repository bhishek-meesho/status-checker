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

	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("❌", link, "might be down!")
		return
	}

	fmt.Println("✅", link, "is up!")
}
