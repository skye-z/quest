# Quest - Exercise Assistant

[中文](README_zh.md)

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

## Experience

[Demo](https://quest.betax.dev)
```
Username: admin
Password: admin
```

> The demo system will be deleted automatically when no one visits for 15 minutes, and deployed automatically when someone visits (very fast, don't worry).

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
git clone https://github.com/skye-z/quest.git
cd quest
bash build.sh
```

Default administrator account: admin / admin

## Firewall pass

If you find that you cannot access other devices after startup, please check whether the firewall is enabled. If so, please pass the port

### Firewall

```shell
firewall-cmd --add-port=80/tcp --permanent
firewall-cmd --reload
```
