package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// 定义互斥锁
	mutex sync.Mutex
	// 使用互斥锁创建条件变量
	cond = sync.NewCond(&mutex)
	// 共享资源
	sharedResource = false
)

// 等待条件满足的 goroutine
func waiter(id int) {
	// 加锁，确保对共享资源的访问是互斥的
	cond.L.Lock()
	// 检查共享资源是否满足条件，如果不满足则等待
	for !sharedResource {
		fmt.Printf("Waiter %d is waiting...\n", id)
		// 调用 Wait 方法会释放锁，并阻塞当前 goroutine，直到被唤醒
		cond.Wait()
	}
	fmt.Printf("Waiter %d is proceeding...\n", id)
	// 解锁
	cond.L.Unlock()
}

// 通知者 goroutine，用于改变条件并通知等待的 goroutine
func notifier() {
	// 模拟一些工作
	time.Sleep(2 * time.Second)
	// 加锁，确保对共享资源的访问是互斥的
	cond.L.Lock()
	// 改变共享资源的状态
	sharedResource = true
	fmt.Println("Notifier is notifying all waiters...")
	// 通知所有等待的 goroutine
	cond.Broadcast()
	// 解锁
	cond.L.Unlock()
}

func main() {
	// 创建多个等待的 goroutine
	for i := 1; i <= 3; i++ {
		go waiter(i)
	}
	// 创建通知者 goroutine
	go notifier()

	// 等待一段时间，确保所有 goroutine 有足够的时间执行
	time.Sleep(5 * time.Second)
	fmt.Println("Main goroutine is exiting.")
}
