# Website Status Checker

A concurrent website monitoring tool built in Go to learn and practice key concurrency concepts.

## 🎯 Project Overview

This project is a learning exercise focused on implementing concurrent programming concepts in Go. It's a website status checker that can monitor multiple websites simultaneously using Go's powerful concurrency features.

## 🔑 Key Learning Concepts

- **Goroutines**: Understanding Go's lightweight threads
- **Channels**: Learning inter-goroutine communication
- **Concurrent Patterns**: Implementing common concurrency patterns
- **Error Handling**: Managing errors in concurrent operations

## 🛠️ Features

- Concurrent website status checking
- Real-time status monitoring
- Support for multiple websites simultaneously
- Status reporting (Up/Down)
- Response time tracking

## 🚀 Technical Implementation

The project demonstrates the use of:
- Multiple goroutines for parallel website checking
- Channels for communication between goroutines
- Context for managing timeouts and cancellation
- Error handling patterns in concurrent code

## 📦 Prerequisites

- Go 1.16 or higher
- Basic understanding of HTTP concepts
- Familiarity with terminal/command line

## 🔧 Installation

```bash
# Clone the repository
git clone [your-repo-url]

# Navigate to project directory
cd status-checker

# Install dependencies (if any)
go mod tidy
```

## 💻 Usage

```bash
# Run the status checker
go run main.go
```

## 📝 Learning Notes

This project serves as a practical implementation of:
- Creating and managing multiple goroutines
- Using channels for data communication
- Implementing concurrent patterns safely
- Managing shared resources in concurrent operations

## 🤝 Contributing

This is a learning project, but suggestions and improvements are welcome! Feel free to:
1. Fork the repository
2. Create a feature branch
3. Submit a pull request

## 📄 License

This project is open source and available under the MIT License.

---
*This is a learning project created while exploring Go's concurrency features.*
