package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"syscall"
)

func main() {
	// 打开文件
	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// 扩展文件大小（确保足够容纳数据）
	fileSize := int64(1024) // 1KB
	if err := file.Truncate(fileSize); err != nil {
		panic(err)
	}

	// 内存映射
	data, err := unix.Mmap(
		int(file.Fd()),                 // 文件描述符
		0,                              // 偏移量
		int(fileSize),                  // 映射长度
		unix.PROT_READ|unix.PROT_WRITE, // 权限：可读可写
		unix.MAP_SHARED,                // 共享映射
	)
	if err != nil {
		panic(err)
	}
	defer func(b []byte) {
		err := unix.Munmap(b)
		if err != nil {
			panic(err)
		}
	}(data) // 解除映射

	// 写入数据
	copy(data, []byte("Hello, mmap!"))

	// 同步到磁盘（确保写入生效）
	if err := unix.Msync(data, syscall.MS_SYNC); err != nil {
		panic(err)
	}

	fmt.Println("Data written via mmap:", string(data))
}
