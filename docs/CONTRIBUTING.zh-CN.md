# 贡献指南

为了方便大家更好地参与到项目中，我们制定了以下的贡献指南。

## 如何开始

1. Fork 本仓库
2. Clone 你 Fork 的仓库到本地
3. 创建一个新的分支，通常会以 `feat/` 开头，表示新功能，`fix/` 表示修复问题
4. 在新的分支上进行开发
5. 提交前可以拉一下最新的原仓库代码，确保没有冲突
6. 提交代码到你 Fork 的仓库
7. 创建一个 Pull Request 到原仓库的 `main` 分支

## 提交规范
- 代码风格：请遵循项目的代码风格，如果你不确定，请参考项目中已有的代码
- 单元测试：请确保你的代码通过了单元测试
- 提交信息：请使用英文书写提交信息，格式为 `type(scope): description`，例如 `feat(client): add new feature`，`fix(server): fix a bug`，`docs: update README.md` 等

## 分支和提交信息
分支和提交信息都需要增加前缀，可以帮助大家快速识别这个提交的类型和范围。

为此我们界定了以下类似：
- type: 表示提交的类型，可以是如下的值：
- feat: 新功能
- fix: 修复问题
- docs: 文档相关的修改
- style: 代码风格相关的修改
- refactor: 重构代码
- test: 测试相关的修改
- chore: 其他的修改
- revert: 撤销之前某次提交

以及如下的范围：
- client: 表示客户端相关的内容
- server：表示服务端相关的内容

通常开发相关的提交信息回需要制定scope，比如 `feat(client): add new feature`，而一些文档或者杂事之类的大多是不需要scope的，比如 `docs: update README.md`。
