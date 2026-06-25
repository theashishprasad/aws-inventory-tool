# AWS Inventory Tool

A Go CLI application that collects AWS infrastructure inventory using the AWS SDK for Go and generates inventory reports across multiple AWS services.

This project is part of my Go learning journey focused on DevOps, Platform Engineering, Cloud Infrastructure, and Automation.

## Project Goal

Build an AWS inventory reporting tool while learning practical Go concepts used in cloud automation and infrastructure tooling.

## Version 4 Features

* Collect EC2 inventory using AWS SDK for Go v2
* Collect Amazon S3 bucket inventory
* Aggregate inventory across multiple AWS services
* Discover AWS credentials using the default credential chain
* Connect to AWS APIs programmatically
* Count EC2 instances across reservations
* Count Amazon S3 buckets
* Generate a formatted inventory report
* Handle AWS SDK errors gracefully
* Organize code using reusable Go packages
* Separate application logic from business logic

## Learning Objectives

### Completed in Version 1

* Structs
* JSON
* Struct Tags
* json.Unmarshal
* File Reading
* Error Handling

### Completed in Version 2

* Packages
* Separation of Concerns
* Exported Functions
* Code Organization
* Reusable Components

### Completed in Version 3

* AWS SDK for Go v2
* Cloud APIs
* AWS Authentication Chain
* config.LoadDefaultConfig()
* EC2 Client Creation
* DescribeInstances API
* API Response Processing
* Infrastructure Inventory Collection

### Completed in Version 4

* Multiple AWS Services
* Amazon S3 SDK
* s3.NewFromConfig()
* ListBuckets API
* Aggregation
* Multi-Service Reporting
* Shared AWS Configuration
* Inventory Consolidation

### Upcoming

* CLI applications
* JSON export
* Concurrency
* Context and timeouts
* Unit testing

## Roadmap

### Version 1

Parse inventory data from a JSON file. ✅

### Version 2

Move logic into reusable packages. ✅

### Version 3

Collect EC2 inventory using AWS SDK. ✅

### Version 4

Add S3 inventory collection. ✅

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
├── inventory/
│   └── inventory.go
├── models/
│   └── inventory.go
├── README.md
├── .gitignore
├── go.mod
└── go.sum
```

## Prerequisites

Install and configure AWS CLI:

```bash
aws configure
```

Verify credentials:

```bash
aws sts get-caller-identity
```

## How to Run

Run the application:

```bash
go run main.go
```

## Sample Output

```text
AWS Inventory Report

EC2 Instances : 12
S3 Buckets    : 8
```

> The actual resource counts depend on the AWS account and region configured on your machine.

## Validation

```bash
go fmt ./...
go mod tidy
go build ./...
go run main.go
```

## Manual Tests

### Valid AWS Credentials

Expected:

```text
AWS Inventory Report

EC2 Instances : X
S3 Buckets    : Y
```

### Missing AWS Credentials

Expected:

```text
Error: ...
```

### Invalid AWS Credentials

Expected:

```text
Error: ...
```

### No EC2 Instances

Expected:

```text
AWS Inventory Report

EC2 Instances : 0
S3 Buckets    : Y
```

### No S3 Buckets

Expected:

```text
AWS Inventory Report

EC2 Instances : X
S3 Buckets    : 0
```

## Technologies Used

* Go
* AWS SDK for Go v2
* Amazon EC2
* Amazon S3
* Cloud APIs
* Packages
* Error Handling

## Learning Outcomes

Through this project I practiced:

* Using the AWS SDK for Go v2
* Authenticating with AWS using the default credential chain
* Creating AWS service clients
* Calling multiple AWS APIs programmatically
* Processing AWS API responses
* Aggregating infrastructure inventory across services
* Creating reusable Go packages
* Separating business logic from application logic
* Handling cloud API errors
* Building infrastructure automation tooling

## Version

Current Version:

```text
v0.4.0
```
