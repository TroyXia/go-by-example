package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// 创建 Redis 客户端
func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "Troy@0403",      // 无密码
		DB:       0,                // 默认数据库
	})

	return client
}

// 演示 Redis Pipeline 工作机制
func main() {
	// 创建 Redis 客户端
	client := createClient()

	// 测试连接
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis: %v", err)
	}

	fmt.Println("成功连接到 Redis 服务器")

	// 清除可能存在的测试数据
	client.Del(ctx, "key1", "key2", "key3")

	// 1. 普通方式执行多个命令
	fmt.Println("\n=== 普通方式执行命令 ===")
	startTime := time.Now()

	client.Set(ctx, "key1", "value1", 0)
	client.Set(ctx, "key2", "value2", 0)
	client.Set(ctx, "key3", "value3", 0)

	val1, _ := client.Get(ctx, "key1").Result()
	val2, _ := client.Get(ctx, "key2").Result()
	val3, _ := client.Get(ctx, "key3").Result()

	fmt.Printf("key1: %s, key2: %s, key3: %s\n", val1, val2, val3)
	fmt.Printf("普通方式耗时: %v\n", time.Since(startTime))

	// 2. 使用 Pipeline 执行多个命令
	fmt.Println("\n=== 使用 Pipeline 执行命令 ===")
	startTime = time.Now()

	// 创建一个 pipeline
	pipe := client.Pipeline()

	// 将多个命令添加到 pipeline
	pipe.Set(ctx, "key1", "pipeline-value1", 0)
	pipe.Set(ctx, "key2", "pipeline-value2", 0)
	pipe.Set(ctx, "key3", "pipeline-value3", 0)

	// 执行 pipeline 中的所有命令
	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Fatalf("Pipeline 执行失败: %v", err)
	}

	// 再次获取值
	val1, _ = client.Get(ctx, "key1").Result()
	val2, _ = client.Get(ctx, "key2").Result()
	val3, _ = client.Get(ctx, "key3").Result()

	fmt.Printf("key1: %s, key2: %s, key3: %s\n", val1, val2, val3)
	fmt.Printf("Pipeline 方式耗时: %v\n", time.Since(startTime))

	// 3. 带事务的 Pipeline
	fmt.Println("\n=== 带事务的 Pipeline ===")
	// 清除数据
	client.Del(ctx, "counter")

	// 创建一个事务性 pipeline
	pipe = client.TxPipeline()

	// 增加计数器
	pipe.Incr(ctx, "counter")
	pipe.Incr(ctx, "counter")
	pipe.Incr(ctx, "counter")

	// 执行事务
	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Fatalf("事务执行失败: %v", err)
	}

	// 获取最终结果
	counter, _ := client.Get(ctx, "counter").Result()
	fmt.Printf("计数器最终值: %s\n", counter)

	// 4. 解释 Redis Pipeline 工作机制
	fmt.Println("\n=== Redis Pipeline 工作机制 ===")
	fmt.Println("1. 普通模式: 客户端发送一个命令，等待服务器响应，然后再发送下一个命令")
	fmt.Println("2. Pipeline 模式: 客户端一次性发送多个命令，不需要等待每个命令的响应")
	fmt.Println("3. 服务器处理: 服务器按顺序执行所有命令，并将所有响应一次性返回给客户端")
	fmt.Println("4. 优势: 减少网络往返时间，提高吞吐量，特别适合执行大量命令的场景")
	fmt.Println("5. 注意事项: Pipeline 中的命令是按顺序执行的，但不保证原子性(除非使用 TxPipeline)")
}
