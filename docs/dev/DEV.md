<div align="center">
  <img src="../logo.png" alt="Logo" width="290" height="251" />
</div>

<div align="center">

| [English](DEV.md) | [中文简体](DEV.zh-CN.md) |

</div>

# Preface

Before starting development, please ensure you have read [README.md](../../README.md) and [GOPHER.md](GOPHER.md). The former helps you understand `iter-x`, while the latter provides programming experience and team guidelines for this project.

# Environment Setup

## 1. protoc Installation

### macOS

```bash
# Install
brew install protobuf
# Verify
protoc --version
```

### Windows

**Command Line Installation**

Open Command Prompt or PowerShell.

```bash
# Install
choco install protoc
# Verify
protoc --version
```

**Manual Installation**

Download the precompiled binary for Windows:
[protobuf releases](https://github.com/protocolbuffers/protobuf/releases)

Extract and add protoc.exe to your system's PATH environment variable.

### Linux

**Ubuntu/Debian:**

```bash
# Install
sudo apt-get update
sudo apt-get install protobuf-compiler
# Verify
protoc --version
```

**Fedora/CentOS/RHEL:**

```bash
# Install
sudo yum install protobuf-compiler
# Verify
protoc --version
```

## 2. Project Environment Dependencies

### Go Version

* Go version >= 1.24.1

### Plugin Installation

```bash
make init
```

### Project Initialization

```bash
make all
```

# Development

This project adopts a mini DDD design philosophy, divided into the following modules:

* API
    * proto
    * pb
* service
* biz
    * bo
    * do
    * repository
* data
    * cache
    * db
    * impl

![go-ddd.png](../images/go-ddd.png)

## Dependency Inversion

The correct calling relationship is:

`service` -> `biz` -> `repository` -> `impl` -> `data`

## Transaction Management

> Transaction management is handled at the biz layer.

* Repository Definition

```bash
// Transaction wrapper interface
type Transaction interface {
	// Exec Transaction execution
	Exec(context.Context, func(ctx context.Context) error) error
}
```