package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个缓冲区
var buffer []int
var mutex sync.Mutex
var cond = sync.NewCond(&mutex)

// 生产者函数
func producer(id int) {
	for {
		// 加锁以保护对缓冲区的操作
		mutex.Lock()
		// 模拟生产数据
		item := id
		buffer = append(buffer, item)
		fmt.Printf("Producer %d produced: %d\n", id, item)
		// 唤醒一个等待的消费者
		cond.Signal()
		// 解锁
		mutex.Unlock()
		// 模拟生产时间
		time.Sleep(1 * time.Second)
	}
}

// 消费者函数
func consumer(id int) {
	for {
		// 加锁以保护对缓冲区的操作
		mutex.Lock()
		// 如果缓冲区为空，则等待
		for len(buffer) == 0 {
			fmt.Printf("Consumer %d is waiting...\n", id)
			cond.Wait()
		}
		// 从缓冲区取出数据
		item := buffer[0]
		buffer = buffer[1:]
		fmt.Printf("Consumer %d consumed: %d\n", id, item)
		// 解锁
		mutex.Unlock()
		// 模拟消费时间
		time.Sleep(2 * time.Second)
	}
}

func main() {
	// 创建生产者和消费者 goroutine
	go producer(1)
	go consumer(1)
	go consumer(2)

	// 让主 goroutine 运行一段时间
	time.Sleep(10 * time.Second)
	fmt.Println("Main goroutine is exiting.")
}
