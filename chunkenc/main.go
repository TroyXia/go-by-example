package main

import (
	"fmt"
	"time"

	"github.com/prometheus/prometheus/tsdb/chunkenc"
)

func main() {
	// 创建一个新的 XOR 压缩块
	chunk := chunkenc.NewXORChunk()

	// 获取块的 Appender 用于添加样本
	appender, err := chunk.Appender()
	if err != nil {
		panic(err)
	}

	now := time.Now().UnixMilli()

	// 添加一些样本数据
	for i := 1; i < 100; i++ {
		timestamp := now + int64(i*1000) // 每秒一个样本
		value := float64(i)
		appender.Append(timestamp, value)
	}

	// 获取块的 Iterator 用于读取样本
	iterator := chunk.Iterator(nil)

	// 读取并打印前10个样本作为示例
	fmt.Println("Timestamp\t\tValue")
	printed := 0
	for iterator.Next() == chunkenc.ValFloat && printed < 10 {
		timestamp, value := iterator.At()
		fmt.Printf("%d\t\t%.2f\n", timestamp, value)
		printed++
	}

	if iterator.Err() != nil {
		panic(iterator.Err())
	}

	// 获取块的一些统计信息
	fmt.Printf("\nChunk stats:\n")
	fmt.Printf("Number of samples: %d\n", chunk.NumSamples())

	// 正确获取压缩后的大小
	chunkBytes := chunk.Bytes()
	fmt.Printf("Compressed size: %d bytes\n", len(chunkBytes))

	// 可选：计算压缩率
	uncompressedSize := chunk.NumSamples() * 16 // 每个样本大约16字节(8字节时间戳+8字节值)
	compressionRatio := float64(uncompressedSize) / float64(len(chunkBytes))
	fmt.Printf("Compression ratio: %.2fx\n", compressionRatio)
}
