package main

import (
	"fmt"

	"github.com/dgraph-io/ristretto"
)

func main() {
	// 初始化Ristretto缓存
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000000, // 最大计数器数量
		MaxCost:     1000000, // 最大成本（可以理解为缓存项的总大小限制）
		BufferItems: 64,      // 缓冲区大小
	})
	if err != nil {
		panic(err)
	}

	// 设置缓存项
	// 设置缓存项 "key1"，值为 "value1"，cost 参数表示该缓存项的成本，用于计算缓存的总大小，此处设置为 1
	cache.Set("key1", "value1", 1)
	cache.Set("key2", 123, 1)
	cache.Set("key3", true, 1)

	// 获取缓存项
	if val, found := cache.Get("key1"); found {
		fmt.Println("key1 的值:", val)
	}

	if val, found := cache.Get("key2"); found {
		fmt.Println("key2 的值:", val)
	}

	if val, found := cache.Get("key3"); found {
		fmt.Println("key3 的值:", val)
	}

	// 不存在的键
	if _, found := cache.Get("not_exist"); !found {
		fmt.Println("not_exist 键不存在")
	}

	// 等待缓存项被处理
	cache.Wait()
}
