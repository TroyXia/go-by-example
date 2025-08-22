# Ristretto 缓存库示例

## 关于 github.com/dgraph-io/ristretto

`github.com/dgraph-io/ristretto` 是由 Dgraph 团队开发的一个高性能、内存高效的键值存储库，特别设计用于构建缓存系统。它的主要特点包括：

- 高命中率
- 快速吞吐量
- 基于成本的逐出策略
- 完全并发访问
- 可选的性能指标

该库灵感来源于 Google 的 Cache 原理论文，旨在解决现代分布式系统中缓存层的挑战。

## 示例代码

本仓库包含两个示例，分别位于不同的目录中：

1. `basic/main.go`: 基本的 Ristretto 缓存使用示例
2. `redis/main.go`: 结合 Redis 和 Ristretto 实现的二级缓存系统

## 如何运行

### 前提条件

- 安装 Go 1.16+ 
- 对于 `redis_ristretto.go`，需要安装并运行 Redis 服务器

### 安装依赖

```bash
# 进入项目目录
cd /Users/hxia/project/go-by-example/ristretto-cache

# 安装依赖（go.mod 文件已提供）
go mod tidy
```

`go.mod` 文件中已包含所需依赖：
- `github.com/dgraph-io/ristretto v0.1.1`
- `github.com/go-redis/redis/v8 v8.11.5`

`go mod tidy` 命令会自动下载并安装这些依赖包。

### 运行示例

```bash
# 运行基本示例
go run basic/main.go

# 运行二级缓存示例
# 确保 Redis 服务器已启动
go run redis/main.go
```

## 示例说明

### main.go

这个示例展示了 Ristretto 缓存的基本用法，包括：
- 初始化缓存
- 设置缓存项
- 获取缓存项
- 检查缓存项是否存在

### redis_ristretto.go

这个示例展示了如何将 Ristretto 与 Redis 结合使用，形成二级缓存系统：
- 一级缓存：本地 Ristretto 缓存，速度快
- 二级缓存：Redis 缓存，可在多个进程间共享
- 数据获取策略：先查本地缓存，再查 Redis，最后查数据库

这种设计可以充分利用本地内存的速度优势，同时借助 Redis 实现数据共享和持久化。