# Games admin server

一个迫真亚运会志愿者管理系统后端

## 部署准备

- nohup 用来后台执行程序
- make 用来调用 makefile

## 如何运行

**运行主程序**

```bash
make run
```

**编译项目**

```bash
make build
```

### 后台服务

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
