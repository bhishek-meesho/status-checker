// Package main implements a website status checker with both
// sequential and concurrent checking options.
package main

import (
	"fmt"
	"net/http"
	"time"
)

// Site represents a website with its URL, status, and response time
type Site struct {
	URL          string
	Status       string
	ResponseTime time.Duration
}

// ANSI color codes for terminal output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
)

// Constants for status messages
const (
	StatusUp   = "UP"
	StatusDown = "DOWN"
)

// checkLink performs an HTTP GET request to check if a website is accessible.
func checkLink(link string) Site {
	// Create a Site object to store the result
	result := Site{URL: link}

	// Record start time
	start := time.Now()

	// Perform the HTTP GET request
	_, err := http.Get(link)

	// Calculate response time
	result.ResponseTime = time.Since(start)

	if err != nil {
		result.Status = StatusDown
		return result
	}

	result.Status = StatusUp
	return result
}

// formatStatus returns colored status string
func formatStatus(status string) string {
	if status == StatusUp {
		return ColorGreen + "✅ " + status + ColorReset
	}
	return ColorRed + "❌ " + status + ColorReset
}

// formatDuration returns colored duration string
func formatDuration(d time.Duration) string {
	if d < 500*time.Millisecond {
		return ColorGreen + d.String() + ColorReset
	} else if d < 1*time.Second {
		return ColorYellow + d.String() + ColorReset
	}
	return ColorRed + d.String() + ColorReset
}

// checkSequentially checks websites one by one
func checkSequentially(sites []string) []Site {
	var results []Site

	fmt.Printf("\n%sChecking websites sequentially...%s\n", ColorCyan, ColorReset)
	for i, site := range sites {
		fmt.Printf("%sChecking (%d/%d) %s...%s", ColorBlue, i+1, len(sites), site, ColorReset)
		result := checkLink(site)
		results = append(results, result)

		// Print immediate result
		fmt.Printf("\r%sChecked  (%d/%d) %s - Status: %s - Response Time: %s%s\n",
			ColorBlue, i+1, len(sites), site,
			formatStatus(result.Status),
			formatDuration(result.ResponseTime),
			ColorReset)
	}

	return results
}

// checkConcurrently checks all websites concurrently using goroutines
func checkConcurrently(sites []string) []Site {
	ch := make(chan Site)
	var results []Site
	checking := make(map[string]bool)

	fmt.Printf("\n%sChecking websites concurrently...%s\n", ColorPurple, ColorReset)

	// Launch a goroutine for each site
	for _, site := range sites {
		checking[site] = true
		go func(url string) {
			ch <- checkLink(url)
		}(site)
	}

	// Collect results from all goroutines
	for range sites {
		result := <-ch
		results = append(results, result)
		delete(checking, result.URL)

		// Print progress
		fmt.Printf("%sReceived result for %s - Status: %s - Response Time: %s%s\n",
			ColorPurple, result.URL,
			formatStatus(result.Status),
			formatDuration(result.ResponseTime),
			ColorReset)

		if len(checking) > 0 {
			fmt.Printf("%sStill waiting for: %v%s\n", ColorYellow, getKeys(checking), ColorReset)
		}
	}

	return results
}

// getKeys returns a slice of keys from a map
func getKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// printResults displays the check results and execution time
func printResults(results []Site, executionTime time.Duration) {
	fmt.Printf("\n%sResults:%s\n", ColorCyan, ColorReset)
	fmt.Printf("%s--------%s\n", ColorCyan, ColorReset)

	var totalTime time.Duration
	var successCount, failCount int

	for _, result := range results {
		fmt.Printf("%-30s Status: %s Response Time: %s\n",
			result.URL,
			formatStatus(result.Status),
			formatDuration(result.ResponseTime))

		totalTime += result.ResponseTime
		if result.Status == StatusUp {
			successCount++
		} else {
			failCount++
		}
	}

	// Print statistics
	fmt.Printf("\n%sStatistics:%s\n", ColorYellow, ColorReset)
	fmt.Printf("Total Websites: %d\n", len(results))
	fmt.Printf("Successful: %s%d%s\n", ColorGreen, successCount, ColorReset)
	fmt.Printf("Failed: %s%d%s\n", ColorRed, failCount, ColorReset)
	fmt.Printf("Average Response Time: %s\n", formatDuration(totalTime/time.Duration(len(results))))
	fmt.Printf("Total Execution Time: %s\n", formatDuration(executionTime))
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

	for {
		// Display menu
		fmt.Printf("\n%s=== Website Status Checker ===%s\n", ColorCyan, ColorReset)
		fmt.Printf("Monitoring %s%d%s websites\n", ColorYellow, len(sites), ColorReset)
		fmt.Printf("\n%sChoose checking method:%s\n", ColorPurple, ColorReset)
		fmt.Printf("%s1. Check one by one (Sequential)%s\n", ColorGreen, ColorReset)
		fmt.Printf("%s2. Check all at once (Concurrent)%s\n", ColorBlue, ColorReset)
		fmt.Printf("%s3. Exit%s\n", ColorRed, ColorReset)
		fmt.Printf("\nEnter your choice (1-3): ")

		// Get user choice
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			startTime := time.Now()
			results := checkSequentially(sites)
			printResults(results, time.Since(startTime))

		case 2:
			startTime := time.Now()
			results := checkConcurrently(sites)
			printResults(results, time.Since(startTime))

		case 3:
			fmt.Printf("\n%sGoodbye!%s\n", ColorCyan, ColorReset)
			return

		default:
			fmt.Printf("\n%sInvalid choice! Please enter 1, 2, or 3.%s\n", ColorRed, ColorReset)
		}
	}
}
