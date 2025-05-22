# Website Status Checker

A Go-based website monitoring tool that demonstrates concurrent programming concepts while providing detailed website status and response time information.

## 🎯 Project Overview

This project is a learning exercise focused on implementing concurrent programming concepts in Go. It's an interactive command-line tool that can check website status either sequentially or concurrently, demonstrating the performance benefits of Go's concurrency features.

## 🛠️ Current Features

- **Interactive Menu System**
  - Sequential checking (one by one)
  - Concurrent checking (all at once)
  - Easy-to-use command-line interface

- **Detailed Status Reporting**
  - Real-time progress updates
  - Color-coded status indicators
  - Individual response times for each website
  - Remaining websites in queue (concurrent mode)

- **Performance Metrics**
  - Individual website response times
  - Color-coded timing indicators:
    - 🟢 Green: < 500ms (fast)
    - 🟡 Yellow: 500ms - 1s (medium)
    - 🔴 Red: > 1s (slow)
  - Total execution time comparison
  - Success/Failure statistics

- **Visual Enhancements**
  - Color-coded output for better readability
  - Progress tracking in both modes
  - Clear statistical summaries
  - Formatted result display

## 💻 Usage

```bash
# Run the status checker
go run main.go

# Menu Options:
1. Check one by one (Sequential)
2. Check all at once (Concurrent)
3. Exit
```

## 📊 Sample Output

```
=== Website Status Checker ===
Monitoring 6 websites

[Sequential Mode]
Checking (1/6) https://www.google.com...
Checked (1/6) https://www.google.com - Status: ✅ UP - Response Time: 245ms

Results:
--------
google.com     Status: ✅ UP    Response Time: 245ms
facebook.com   Status: ✅ UP    Response Time: 612ms

Statistics:
Total Websites: 6
Successful: 6
Failed: 0
Average Response Time: 428ms
Total Execution Time: 1.2s
```

## 🔑 Key Learning Concepts

- **Goroutines**: Implementation of concurrent website checking
- **Channels**: Communication between goroutines for result collection
- **Error Handling**: Proper HTTP request error management
- **Time Measurement**: Response time tracking and formatting
- **ANSI Colors**: Terminal output enhancement
- **Structured Data**: Using custom types for organized data handling

## 🚀 Technical Implementation

The project demonstrates:
- Concurrent vs Sequential execution patterns
- Channel usage for goroutine communication
- Custom type definitions for structured data
- Terminal output formatting and coloring
- Time measurement and formatting
- Error handling in network requests

## 📦 Prerequisites

- Go 1.16 or higher
- Terminal with ANSI color support

## 🤝 Contributing

This is a learning project, but suggestions and improvements are welcome! Feel free to:
1. Fork the repository
2. Create a feature branch
3. Submit a pull request

## 📄 License

This project is open source and available under the MIT License.

---
*This is a learning project created while exploring Go's concurrency features and terminal-based user interfaces.*
