package main

import (
	"context"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hxia/.kube/config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	podInformer := v1.NewPodInformer(
		clientset,
		metav1.NamespaceAll,
		1*time.Minute,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*apicorev1.Pod)
			log.Printf("Pod created: %s in namespace %s", pod.Name, pod.Namespace)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			newPod := newObj.(*apicorev1.Pod)
			log.Printf("Pod updated: %s in namespace %s", newPod.Name, newPod.Namespace)
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*apicorev1.Pod)
			log.Printf("Pod deleted: %s in namespace %s", pod.Name, pod.Namespace)
		},
	})

	stopCh := context.TODO().Done()
	go podInformer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, podInformer.HasSynced) {
		log.Fatalf("Failed to sync podInformer")
	}

	log.Println("Starting to watch for pod events")
	select {}
}
