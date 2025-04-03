# Standard Operating Procedures (SOP) Manual

服务器端语言是Go，入口在 `cmd/server` 目录下，我们采用了 `github.com/google/wire` 进行依赖注入，所以如果你有新的服务需要到 `main.go` 和 `wire.go` 中看看是否需要新增。

整体分为几个层：
1. `server`：包括gRPC和HTTP入口（HTTP-Gateway最终也是落到gRPC的）
2. `service`：gRPC之后的入口，主要做一些参数校验，参数从gRPC对象转成bo对象
3. `biz`：业务逻辑层，主要是业务逻辑处理，比如调用外部服务、数据处理等，大部分逻辑都集中在这里
4. `data`：数据层，主要是数据的读写，包括数据库、外部服务等，这里接受 `bo` 对象（在 `internal/biz/bo/*.go`），但是返回的话需要返回 `do` 对象（在 `internal/biz/do/*.go`），数据库（目前ent）和外部的服务的对象独立不可透传回 `biz` 和 `service` 层

## Makefile
项目大量地方涉及自动生成代码的应用，所有涉及的命令都在 `Makefile` 里了，比如：
- `make all`: 生成所有的代码
- `make config`: 生成配置
- `make api` : 从 `proto/**/*.proto` 生成对应的 `*.go` 文件

## 配置
配置在 `config` 目录下，主要是 `*.yaml` 文件，配置里的一些敏感的是需要定义在 `.env.example` 里的。我们是用 `proto/conf/conf.proto` 定义的配置，然后通过 `make config` 生成代码的

## API
全新创建一个API需要涉及这么几个步骤：
1. 到 `proto/api/module/*.proto` 中定义
2. 然后运行 `make api` 生成对应的代码
3. 在 `server` 层注册对应的服务
4. 在 `service` 层实现对应的方法
5. 在 `biz` 层实现对应的逻辑
6. 在 `data` 层实现对应的数据操作（有可能原来就有存在的，需要自己判断）

## 错误
错误码在 `proto/xerr/xerr.proto` 中定义，然后通过 `make errors && make generate` 生成对应的错误码，我们通过 `i18n/*.toml` 来进行错误码的国际化，所以有新的错误生成的时候需要补充对应语言的错误提示。在 `biz` 层应该拦截所有的错误，并返回对应的 `xerr` 错误给到 `service` 层，这样避免把错误信息暴露给到客户端

## 数据库
目前我们使用的是 `Postgres` ，Go里用的是 `ent` 这个ORM，所有的数据操作都是通过 `ent` 这个库来进行的，所以如果有新的表需要创建，需要在 `internal/data/ent/schema/` 下创建对应的 `*.go` 文件，然后运行 `make all` 生成对应的代码，然后在 `data` 层进行对应的实现。

schema里有一个 `mixin.go` 有需要的话可以用

## 测试
每次完成所有任务后请执行一下 `make all` ，然后尝试编译一下 `make build` ，确保都没有问题，如果有问题要解决

## 代码注释
通常不需要代码注释，更多是用规范的命令来表现意图，不过有一些地方必须要有注释的请用英文注释
