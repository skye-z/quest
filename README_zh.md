# Quest - 在线题库与模拟考试

[English](README.md)

Quest 是一个用于题库管理和模拟考试的超小型系统.

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

## 在线体验

[演示系统](https://quest.betax.dev)
```
账户: admin
密码: admin
```

> 演示系统在15分钟无人访问时自动删除, 在有人访问时自动部署(耗时10秒内).

## 最低硬件要求

单核、16MB内存、64MB存储空间即可.

Quest 对硬件要求很低, 您几乎可以在任何设备上运行它.

## 基准测试

下面使用查询科目列表API(需登录)作为测试接口, 每轮测试循环10次.

测试设备: Mac Mini 2018 (i3 4core 3.6GHz / 20G DDR4 2666)

测试软件: JMeter 5.5

| Thread | 90% | 95% | 99% | Error | QPS | CPU | Memory |
| ------ | --- | --- | --- | ----- | --- | --- | --- |
| 5000 | 2240 | 2875 | 3902 | 0.02% | 3230.4/s | 191.2% | 659MB |
| 3000 | 1255 | 1618 | 2242 | 0.09% | 4042.9/s | 172.3% | 275MB |
| 1000 | 505 | 619 | 893 | 0.00% | 3125.0/s | 145.8% | 154MB |

根据测试结果, 虽然您可以在任何设备上运行它, 但如果您想获得3000QPS的速度, 最好使用2核以上的CPU并提供1G内存.

## 从头构建
```shell
git clone https://github.com/skye-z/quest.git
cd quest
bash build.sh
```

默认管理员账户: admin / admin

## 防火墙放行

如果您发现启动后无法在外部访问, 请检查防火墙是否启动, 如启动则请放行80端口

### Firewall

```shell
firewall-cmd --add-port=80/tcp --permanent
firewall-cmd --reload
```
