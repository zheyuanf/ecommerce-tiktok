## 概述
本项目为 Marscode 训练营后端方向项目，由团队 e706队 开发

## 代码结构描述
本项目的代码结构主要分为以下几个部分：

- `app/`: 包含了项目的应用代码，包括各个服务的实现。
- `rpc_gen/`: 包含了使用 `cwgo` 生成的 RPC 客户端代码。
- `idl/`: 包含了项目使用的接口定义语言（IDL）文件。
- `scripts/`: 包含了项目中使用的脚本文件，例如代码生成脚本。
- `go.work`: 定义了项目的 Go 模块结构。

## Makefile 命令

### 一般命令
- **help**: 显示帮助信息。
  ```bash
  make help
  ```

### 代码生成
- **gen**: 为指定服务生成客户端代码。示例用法：
  ```bash
  make gen svc='product'
  ```

- **gen-client**: 使用 `cwgo` 为指定服务生成客户端代码。示例用法：
  ```bash
  make gen-client svc='product'
  ```

- **gen-server**: 使用 `cwgo` 为指定服务生成服务器代码。示例用法：
  ```bash
  make gen-server svc='product'
  ```

### 开发环境
- **env-start**: 启动所有中间件软件，使用 Docker。
  ```bash
  make env-start
  ```

- **env-stop**: 停止所有正在运行的 Docker 容器。
  ```bash
  make env-stop
  ```

### 清理
- **clean**: 清理在构建过程中生成的所有临时文件。
  ```bash
  make clean
  ```

## 使用方法
要使用 Makefile 命令，只需在终端中运行 `make <command>`，将 `<command>` 替换为上述列出的命令之一。

## 需求
确保您已安装必要的工具，例如 `Docker` 和 `cwgo`，以成功执行命令。

## 贡献
欢迎通过提交问题或拉取请求来为本项目做出贡献。

## 贡献
欢迎通过提交问题或拉取请求来为本项目做出贡献。

## 贡献
欢迎通过提交问题或拉取请求来为本项目做出贡献。