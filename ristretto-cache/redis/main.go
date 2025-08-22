package main

import (
	"context"
	"fmt"

	"github.com/dgraph-io/ristretto"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// 初始化Ristretto缓存（一级缓存）
	localCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000000,
		MaxCost:     1000000,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}

	// 初始化Redis客户端（二级缓存）
	// 如果Redis服务器设置了密码，请在下面的Password字段中设置
	// 例如: Password: "your_redis_password",
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "Troy@0403",      // Redis密码，如果没有设置密码则为空字符串
		// DB:       0,                // 使用的数据库索引，默认为0
	})

	// 测试Redis连接
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis连接失败:", err)
		// 注意：如果没有Redis，程序仍可以使用本地缓存
	}

	// 模拟从数据库获取数据的函数

	// 获取数据的函数（先查本地缓存，再查Redis，最后查数据库）
	getData := func(key string) string {

		// 先将数据写入Redis和缓存
		data := "data for " + key
		localCache.Set(key, data, int64(len(data)))
		if redisClient != nil {
			redisClient.Set(ctx, key, data, 0) // 0表示永不过期
		}
		localCache.Wait() // 等待缓存写入完成

		// 1. 尝试从本地缓存获取
		if val, found := localCache.Get(key); found {
			fmt.Println("从本地缓存获取数据:", key)
			return val.(string)
		}

		// 2. 尝试从Redis获取
		if redisClient != nil {
			val, err := redisClient.Get(ctx, key).Result()
			if err == nil {
				fmt.Println("从Redis获取数据:", key)
				// 更新本地缓存
				localCache.Set(key, val, int64(len(val)))
				return val
			}
		}

		return ""
	}

	// 测试
	fmt.Println("第一次调用getData('test_key'):")
	data1 := getData("test_key")
	fmt.Println("结果:", data1)

	fmt.Println("\n第二次调用getData('test_key'):")
	data2 := getData("test_key")
	fmt.Println("结果:", data2)

	// 等待缓存项被处理
	localCache.Wait()
}
