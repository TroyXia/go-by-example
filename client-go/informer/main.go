package main

import (
	"log"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 解析 kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hxia/.kube/config")
	if err != nil {
		panic(err)
	}

	// 创建 ClientSet 客户端对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	// 创建 sharedInformers
	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute)
	// 创建 informer
	informer := sharedInformers.Core().V1().Pods().Informer()

	// 创建 Event 回调 handler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("New Pod Added to Store: %s", mObj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(v1.Object)
			nObj := newObj.(v1.Object)
			log.Printf("%s Pod Updated to %s", oObj.GetName(), nObj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("Pod Deleted from Store: %s", mObj.GetName())
		},
	})

	// 运行 informer
	informer.Run(stopCh)
}
