# Quest - Exercise Assistant

Quest is a smart software for question bank management and mock exams. 

[![](https://img.shields.io/badge/Go-1.20+-%2300ADD8?style=flat&logo=go)](go.work)
[![](https://img.shields.io/badge/Quest-1.0.0-green)](control)

## Features
1. Allows users to import their own question banks and then generate test papers from them for mock exams;
2. Automatically adjust the difficulty and scope of the test papers according to the user’s preferences and level;
3. Record the user’s answers and scores, provide detailed analysis and suggestions.

## Compile and package
```shell
go build -o quest -ldflags '-s -w'

# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o quest -ldflags '-s -w'
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o quest -ldflags '-s -w'
```