package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

// 模拟一个需要等待的条件
func conditionToWait() (bool, error) {
	// 这里可以是任何需要检查的条件，例如检查某个资源是否就绪
	// 为了演示，我们简单地模拟一个条件
	// 假设我们等待一个计数器达到 5
	staticCounter := 0
	return func() (bool, error) {
		staticCounter++
		fmt.Printf("当前计数器值: %d\n", staticCounter)
		if staticCounter >= 5 {
			return true, nil
		}
		return false, nil
	}()
}

var staticCounter int

func conditionToWait2() (bool, error) {
	// 这里可以是任何需要检查的条件，例如检查某个资源是否就绪
	// 为了演示，我们简单地模拟一个条件
	// 假设我们等待一个计数器达到 5
	staticCounter++
	fmt.Printf("当前计数器值: %d\n", staticCounter)
	if staticCounter >= 5 {
		return true, nil
	}
	return false, nil
}

func main() {
	interval := 1 * time.Second
	timeout := 10 * time.Second

	err := wait.Poll(interval, timeout, conditionToWait)
	if err != nil {
		if err == wait.ErrWaitTimeout {
			fmt.Println("等待超时，条件未满足")
		} else {
			fmt.Printf("等待过程中出现错误: %v\n", err)
		}
	} else {
		fmt.Println("条件已满足")
	}

	err = wait.Poll(interval, timeout, conditionToWait2)
	if err != nil {
		if err == wait.ErrWaitTimeout {
			fmt.Println("等待超时，条件未满足")
		} else {
			fmt.Printf("等待过程中出现错误: %v\n", err)
		}
	} else {
		fmt.Println("条件已满足")
	}
}
