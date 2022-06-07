# Games admin server 🏂

一个迫真亚运会志愿者管理系统后端 😝

[![Go](https://github.com/dump-time/games-admin-server/actions/workflows/go.yml/badge.svg)](https://github.com/dump-time/games-admin-server/actions/workflows/go.yml)

## 部署准备 🏋️

- nohup 用来后台执行程序
- make 用来调用 makefile

### 管理员账号 👽

自行通过 navicat 插入或者导入管理员数据信息

## 如何运行 🚀

**运行主程序**

```bash
make run
```

**编译项目**

```bash
make build
```

### 配置文件 ⚙️

本程序通过 --config 参数指定配置文件路径

项目 data 目录下有一个 config.example.yml 的文件，是配置文件的模板。

### 管理员数据 👮

本程序通过 --excel 参数指定管理员数据excel表的路径

项目 data 目录下有一个 team_admins.example.xlsx 的文件，是管理员信息文件的模板。

### 后台服务 🐸

> %%%, 服务稳定++, 运行时长续 1s

**启动服务**

启动服务会自动判断是否有服务在运行，如果有，那么就会先停止服务，再运行

启动服务前会自动编译

```bash
make start
```

**停止服务**

```bash
make stop
```
