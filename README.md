# Quest - Exercise Assistant

Quest is a ultra mini system for question bank management and mock exams. 

[![](https://img.shields.io/badge/Go-1.20+-%2300ADD8?style=flat&logo=go)](go.work)
[![](https://img.shields.io/badge/Quest%20Service-1.0.1-green)](control)
[![](https://img.shields.io/badge/Quest%20Page-1.0.1-blue)](https://github.com/skye-z/quest-page)
[![](https://img.shields.io/badge/Quest%20Extension-1.0.0-red)](https://github.com/skye-z/quest-extension)
[![](https://img.shields.io/badge/License-GPL%20v3.0-orange)](LICENSE)

[![CodeQL](https://github.com/skye-z/quest/workflows/CodeQL/badge.svg)](https://github.com/skye-z/quest/security/code-scanning)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_quest&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_quest)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_quest&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_quest)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_quest&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_quest)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=skye-z_quest&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=skye-z_quest)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=skye-z_quest&metric=bugs)](https://sonarcloud.io/summary/new_code?id=skye-z_quest)

## TODO

* [x] API
    * [x] User
    * [x] System
    * [x] Subject
    * [x] Question
    * [x] Exam
* [x] Page
    * [x] Auth
    * [x] Home
    * [x] Admin
    * [x] Question
    * [x] Search
    * [x] Exam
* [x] Extension
    * [x] Auth
    * [x] Inject
    * [x] Scan
    * [x] Push

## Features
1. Allows users to import their own question banks and then generate test papers from them for mock exams;
2. Automatically adjust the difficulty and scope of the test papers according to the user’s preferences and level;
3. Record the user’s answers and scores, provide detailed analysis and suggestions.

## Minimum System Requirements

Single core, 16MB RAM, 64MB storage is sufficient

This is a very low requirement for modern computers, You can run it on almost any device

## Benchmark

The following uses the query subject list api (login required) as the test interface, and each test cycles 10 times

Test PC: Mac Mini 2018 (i3 4core 3.6GHz / 20G DDR4 2666)

Test software: JMeter 5.5

| Thread | 90% | 95% | 99% | Error | QPS | CPU | Memory |
| ------ | --- | --- | --- | ----- | --- | --- | --- |
| 5000 | 2240 | 2875 | 3902 | 0.02% | 3230.4/s | 191.2% | 659MB |
| 3000 | 1255 | 1618 | 2242 | 0.09% | 4042.9/s | 172.3% | 275MB |
| 1000 | 505 | 619 | 893 | 0.00% | 3125.0/s | 145.8% | 154MB |

According to the test results, although you can run it on any device, if you want to get 3000QPS you'd better use a cpu with more than 2 cores and provide 1g of RAM

## Compile and package
```shell
go mod download
go mod tidy

go build -o quest -ldflags '-s -w'

# MacOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o quest -ldflags '-s -w'
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o quest -ldflags '-s -w'
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o quest -ldflags '-s -w'
```

Default administrator account: admin / admin

## Firewall pass

If you find that you cannot access other devices after startup, please check whether the firewall is enabled. If so, please pass the port

### Firewall

```shell
firewall-cmd --add-port=80/tcp --permanent
firewall-cmd --reload
```
