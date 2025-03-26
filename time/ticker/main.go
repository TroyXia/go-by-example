package main

import (
	"fmt"
	"time"

	"k8s.io/utils/clock"
)

func main() {
	// 创建一个实际的时钟实例
	realClock := clock.RealClock{}

	// 设置定时器的延迟时间为 5 秒
	delay := 5 * time.Second

	// 使用时钟创建一个定时器
	timer := realClock.NewTimer(delay)

	fmt.Println("定时器已启动，等待 5 秒...")

	// 等待定时器触发
	<-timer.C()
	fmt.Println("定时器触发，2 秒已过。")

	// 重置定时器，再次设置延迟时间为 1 秒
	timer.Reset(1 * time.Second)
	fmt.Println("定时器已重置，等待 1 秒...")

	// 再次等待定时器触发
	<-timer.C()
	fmt.Println("定时器再次触发，1 秒已过。")

	// 停止定时器
	timer.Stop()
	fmt.Println("定时器已停止。")

	// 再次等待定时器触发
	timer.Reset(3 * time.Second)
	fmt.Println("重置定时器，等待 3 秒...")
	ticker := timer.C()

	jacker := make(chan string)

	go func() {
		for {
			jacker <- "ma"
			time.Sleep(1 * time.Second)
		}
	}()

	stopCh := make(chan int)
	go func() {
		for {
			select {
			case <-ticker:
				fmt.Println("定时器再次触发，3 秒已过。")
				stopCh <- 1
			case <-jacker:
				fmt.Println("I am jacker")
			}
		}
	}()

	<-stopCh
	fmt.Println("stopped")
}
