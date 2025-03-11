package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载 kubeconfig 文件
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Fatalf("Failed to get kubeconfig: %v", err)
	}

	// 创建 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create clientset: %v", err)
	}

	// 设置 limit 参数
	limit := int64(1)
	// 初始化 continue 字段
	continueToken := ""

	for {
		// 创建 ListOptions
		listOptions := metav1.ListOptions{
			Limit:    limit,
			Continue: continueToken,
		}

		// 获取 Pod 列表
		pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), listOptions)
		if err != nil {
			log.Fatalf("Failed to list pods: %v", err)
		}

		// 打印当前页的 Pod 名称
		for _, pod := range pods.Items {
			fmt.Printf("Pod Name: %s, continue id: %s\n", pod.Name, listOptions.Continue)
		}

		// 获取 continue 字段
		continueToken = pods.Continue
		if continueToken == "" {
			// 没有更多数据，退出循环
			break
		}
	}
}
