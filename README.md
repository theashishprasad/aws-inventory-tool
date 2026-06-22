# AWS Inventory Tool

A Go CLI application that reads AWS inventory data from a JSON file and generates a simple inventory report.

This project is part of my Go learning journey focused on DevOps, Platform Engineering, Cloud Infrastructure, and Automation.

## Project Goal

Build an AWS inventory reporting tool while learning practical Go concepts used in cloud automation and infrastructure tooling.

## Version 1 Features

* Read inventory data from a JSON file
* Parse JSON into Go structs
* Generate a formatted inventory report
* Handle file reading errors gracefully
* Handle JSON parsing errors gracefully

## Learning Objectives

### Completed in Version 1

* Structs
* JSON
* Struct Tags
* json.Unmarshal
* File Reading
* Error Handling

### Upcoming

* Packages and project organization
* AWS SDK for Go
* CLI applications
* JSON export
* Concurrency
* Context and timeouts
* Unit testing

## Roadmap

### Version 1

Parse inventory data from a JSON file. ✅

### Version 2

Move logic into reusable packages.

### Version 3

Collect EC2 inventory using AWS SDK.

### Version 4

Add S3 inventory collection.

### Version 5

Accept AWS region as command-line input.

### Version 6

Export inventory as JSON.

### Version 7

Collect inventory concurrently.

### Version 8

Add context and timeout handling.

### Version 9

Add unit tests.

### Version 10

Production-ready inventory reporting tool.

## Project Structure

```text
aws-inventory-tool/
├── main.go
├── models/
│   └── inventory.go
├── sample/
│   └── inventory.json
├── README.md
├── .gitignore
└── go.mod
```

## How to Run

Run the application:

```bash
go run main.go
```

## Sample Input

```json
{
  "instance_count": 12,
  "region": "ap-south-1"
}
```

## Sample Output

```text
AWS Inventory Report

Instances : 12
Region    : ap-south-1
```

## Validation

```bash
go fmt ./...
go build ./...
go run main.go
```

## Manual Tests

### Valid JSON

Expected:

```text
AWS Inventory Report

Instances : 12
Region    : ap-south-1
```

### Missing File

Expected:

```text
Error: ...
```

### Invalid JSON

Expected:

```text
Error: ...
```

## Technologies Used

* Go
* JSON
* Structs
* File I/O
* Error Handling

## Learning Outcomes

Through this project I practiced:

* Reading files in Go
* Parsing JSON using json.Unmarshal
* Working with structs and struct tags
* Handling errors using idiomatic Go patterns
* Building a simple CLI application
* Mapping JSON data to Go data structures

## Version

Current Version:

```text
v0.1.0
```
