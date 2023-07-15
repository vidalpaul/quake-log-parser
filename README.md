# Quake Log Parser

This project is a Quake log parser implemented in Go, developed specifically for solving [this assignment](https://gist.github.com/cloudwalk-tests/704a555a0fe475ae0284ad9088e203f1). It allows you to read a Quake log file, extract game data for each match, collect kill data, and generate reports based on the parsed information.

## Requirements

To use this project, you need the following:

- Go programming language
- Git

## Getting Started

Follow these steps to get started with the Quake log parser:

1. Clone the repository:

   ```bash
   git clone github.com/vidalpaul/cloudwalk-quake-test.git
   ```

2. Build the project:

   ```bash
   go build
   ```

3. Run the project:

   ```bash
   go run main.go
   ```

4. Run tests:

   ```bash
   go test ./...
   ```

Alternatively, you can run the project using Docker:

```bash
docker build -t quake-log-parser -f build/Dockerfile .

docker run --rm quake-log-parser
```
