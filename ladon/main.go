// 基于ladon的授权示例
package main

import (
	"fmt"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
)

func main() {
	// 初始化策略管理器（使用内存存储）
	warden := &ladon.Ladon{
		Manager: memory.NewMemoryManager(),
		// 移除非标准字段初始化
	}

	// 定义访问控制策略
	policy := &ladon.DefaultPolicy{
		ID:          "user-access-policy",
		Description: "允许用户访问自己的资源",
		Subjects:    []string{"user:<.*>"}, // 主体匹配正则
		Effect:      ladon.AllowAccess,     // 允许访问
		Resources:   []string{"resource:users:<.*>"}, // 资源匹配正则
		Actions:     []string{"get", "update"},      // 允许的操作
		Conditions: ladon.Conditions{
			"owner": &ladon.StringEqualCondition{
				Equals: "alice", // 直接匹配用户ID
			},
		},
	}

	// 添加策略到管理器
	if err := warden.Manager.Create(policy); err != nil {
		panic(err)
	}

	// 构造授权请求
	request := &ladon.Request{
		Subject:  "user:alice",      // 当前用户
		Action:   "get",             // 尝试的操作
		Resource: "resource:users:alice", // 访问的资源
		Context: ladon.Context{
			"owner": "alice", // 资源所有者
		},
	}

	// 执行权限检查
	if err := warden.IsAllowed(request); err != nil {
		fmt.Printf("拒绝访问理由: %+v\n", err)
	} else {
		fmt.Println("允许访问")
	}
}