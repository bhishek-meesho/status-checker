// Package main implements a concurrent website status checker
// that demonstrates Go's concurrency patterns using goroutines and channels.
package main

import (
	"fmt"
	"net/http"
	"time"
)

// Site represents a website with its URL and current status
type Site struct {
	URL    string
	Status string
}

// Constants for status messages
const (
	StatusUp   = "✅ UP"
	StatusDown = "❌ DOWN"
)

// checkLink performs an HTTP GET request to check if a website is accessible.
// It sends the result back through the provided channel.
//
// Parameters:
//   - link: The URL to check
//   - ch: Channel to send the check result
//
// The function demonstrates basic error handling and channel communication.
func checkLink(link string, ch chan Site) {
	// Perform the HTTP GET request
	_, err := http.Get(link)

	// Create a Site object to store the result
	result := Site{URL: link}

	if err != nil {
		result.Status = StatusDown
		ch <- result
		return
	}

	result.Status = StatusUp
	ch <- result
}

// monitorSites initiates concurrent monitoring of multiple websites
// and returns a map of their statuses.
//
// Parameters:
//   - sites: Slice of website URLs to monitor
//
// Returns:
//   - map[string]string: Map of website URLs to their current status
func monitorSites(sites []string) map[string]string {
	// Create a channel to receive results
	// The channel type is Site to handle structured data
	ch := make(chan Site)

	// Map to store the final status of each site
	status := make(map[string]string)

	// Launch a goroutine for each site
	for _, site := range sites {
		go checkLink(site, ch)
	}

	// Collect results from all goroutines
	for range sites {
		result := <-ch
		status[result.URL] = result.Status
	}

	return status
}

func main() {
	// List of websites to monitor
	sites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
	}

	// Print start message
	fmt.Println("Starting website status checker...")
	fmt.Printf("Monitoring %d sites...\n\n", len(sites))

	// Record start time for performance measurement
	startTime := time.Now()

	// Monitor sites concurrently and get results
	results := monitorSites(sites)

	// Print results in a formatted way
	fmt.Println("Results:")
	fmt.Println("--------")
	for url, status := range results {
		fmt.Printf("%-30s %s\n", url, status)
	}

	// Print execution time
	fmt.Printf("\nTotal execution time: %v\n", time.Since(startTime))
}
