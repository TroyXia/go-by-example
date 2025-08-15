# Go Redis 示例

这个仓库包含了使用 Go 语言操作 Redis 的示例代码。

## Pub/Sub 示例

`pubsub` 目录下包含了一个实现 Redis 发布订阅功能的示例。

### 功能说明

这个示例演示了如何使用 Go 语言连接 Redis 服务器，并实现发布订阅功能：
1. 连接到 Redis 服务器
2. 创建一个订阅者监听指定频道
3. 创建一个发布者向指定频道发送消息
4. 订阅者接收并打印消息

### 运行步骤

1. 确保你已经安装了 Redis 服务器并正在运行

2. 安装依赖
```bash
cd /Users/hxia/project/go-by-example/redis
go mod tidy
```

3. 运行示例
```bash
cd pubsub
go run main.go
```

### 代码说明

- `main.go`: 主程序文件，包含以下主要部分：
  - `createClient()`: 创建 Redis 客户端连接
  - `publish()`: 发布消息到指定频道
  - `subscribe()`: 订阅指定频道并接收消息
  - `main()`: 主函数，协调发布和订阅操作

## 依赖

- Go 1.21 或更高版本
- Redis 服务器
- go-redis 客户端库