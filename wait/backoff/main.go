package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/utils/clock"
	"time"
)

// 自定义 BackoffManager 实现
type CustomBackoffManager struct {
	backoff wait.Backoff
	clock   clock.Clock
}

// Backoff 方法实现 BackoffManager 接口
func (c *CustomBackoffManager) Backoff() clock.Timer {
	duration := c.backoff.Step()
	return c.clock.NewTimer(duration)
}

// 模拟一个可能失败的操作，这里随机返回操作是否成功
func tryOperation() bool {
	// 模拟 50% 的失败概率
	success := time.Now().UnixNano()%2 == 0
	if success {
		fmt.Println("Operation succeeded!")
	} else {
		fmt.Println("Operation failed, retrying...")
	}
	return false
}

func main() {
	// 定义指数退避策略
	backoff := wait.Backoff{
		// 初始等待时间
		Duration: 500 * time.Millisecond,
		// 指数因子，每次重试等待时间会乘以该因子
		Factor: 2,
		// 随机抖动因子，避免多个重试任务同时执行
		Jitter: 0.1,
		// 最大重试次数
		Steps: 5,
		// 最大等待时间
		Cap: 5 * time.Second,
	}

	// 创建自定义 BackoffManager 实例
	customBackoffManager := &CustomBackoffManager{
		backoff: backoff,
		clock:   clock.RealClock{},
	}

	// 定义停止信号通道
	stopCh := make(chan struct{})
	go func() {
		// 模拟 10 秒后停止重试
		time.Sleep(10 * time.Second)
		close(stopCh)
	}()

	// 使用 wait.BackoffUntil 执行操作
	wait.BackoffUntil(func() {
		if tryOperation() {
			// 如果操作成功，关闭停止信号通道，结束重试
			close(stopCh)
		}
	}, customBackoffManager, true, stopCh)

	fmt.Println("Retry process ended.")
}
