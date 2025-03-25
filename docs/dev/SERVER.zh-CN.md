<div align="center">
  <img src="../logo.png" alt="Logo" width="290" height="251" />
</div>

<div align="center">

| [English](SERVER.md) | [中文简体](SERVER.zh-CN.md) |

</div>

# 写在前面

在开发前，请确保你已经阅读了[README.md](../README.zh-CN.md), 以及[GOPHER.md](./GOPHER.zh-CN.md), 前者帮助你了解`Iter-X`, 后者给你一些关于此项目的编程经验和团队规范

# 环境篇

## 1. protoc 安装

### MacOS

```bash
# 安装
brew install protobuf
# 验证
protoc --version
```

### Windows

**命令行安装**

打开命令提示符或 PowerShell。

```bash
# 安装
choco install protoc
# 验证
protoc --version
```

**手动安装**

下载适用于 Windows 的预编译二进制文件：
[protobuf releases](https://github.com/protocolbuffers/protobuf/releases)

解压并将 protoc.exe 添加到系统环境变量 PATH 中。

### Linux

**Ubuntu/Debian：**

```bash
# 安装
sudo apt-get update
sudo apt-get install protobuf-compiler
# 验证
protoc --version
```

**Fedora/CentOS/RHEL：**

```bash
# 安装
sudo yum install protobuf-compiler
# 验证
protoc --version
```

## 2. 项目环境依赖

### Go版本

* Go版本 >= 1.24.1

### 相关插件安装

```bash
make init
```

### 项目初始化

```bash
make all
```

# 开发篇

本项目采用mini [DDD](https://www.google.com/search?q=DDD) 设计思想，总共分为以下几个模块

* API: 接口层，包含了HTTP、gRPC等接口定义
  * proto
  * pb
* service: 服务入口，承接了API(Server)和biz层，通常做一些参数校验、转换等工作
* biz: 业务逻辑层，包含了业务逻辑的实现
  * bo
  * do
  * repository
* data: 数据访问层，负责包括数据库、外部数据和外部服务等资源的处理
  * cache
  * db
  * impl

![go-ddd.png](../images/go-ddd.zh-CN.png)

## 依赖倒置

正确的调用关系是：

`service` -> `biz` -> `repository` -> `impl` -> `data`

需要注意的是，每一层都会涉及到数据结构的转置，也就是每层只关心自己收到的数据结构，如：
- `service` 层只关心 `pb` 包中的数据结构，也就是 `proto` 生成的 `req` 和 `resp`，当 `service` 调用 `biz` 时，会将 `pb` 转成 `bo`

## 事务管理

> biz 层使用事务管理

* repository 定义

```go
// Transaction wrapper interface
type Transaction interface {
	// Exec Transaction execution
	Exec(context.Context, func(ctx context.Context) error) error
}
```
