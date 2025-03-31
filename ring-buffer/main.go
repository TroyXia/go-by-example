package main

import (
	"fmt"
	"k8s.io/utils/buffer"
)

func main() {
	// 创建一个初始容量为 2 的可增长环形缓冲区
	ring := buffer.NewRingGrowing(2)

	// 向缓冲区写入数据
	dataToWrite := []int{1, 2, 3, 4, 5}
	for _, num := range dataToWrite {
		ring.WriteOne(num)
	}

	// 从缓冲区读取数据
	for {
		item, ok := ring.ReadOne()
		if !ok {
			break
		}
		fmt.Println(item)
	}
}
