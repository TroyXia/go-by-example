package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"strings"
)

func UsersIndexFunc(obj interface{}) ([]string, error) {
	pod := obj.(*v1.Pod)
	usersString := pod.Annotations["users"]

	return strings.Split(usersString, ","), nil
}

func main() {
	index := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{
		//"byUser":    UsersIndexFunc,
		"namespace": cache.MetaNamespaceIndexFunc,
	})

	pod1 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "one", Namespace: "default",
		Annotations: map[string]string{"users": "ernie,bert"}}}
	pod2 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "two", Namespace: "default",
		Annotations: map[string]string{"users": "bert,oscar"}}}
	pod3 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "three", Namespace: "default",
		Annotations: map[string]string{"users": "ernie,elmo"}}}

	index.Add(pod1)
	index.Add(pod2)
	index.Add(pod3)

	/*
		erniePods, err := index.ByIndex("byUser", "ernie")
		if err != nil {
			panic(err)
		}

		for _, erniePod := range erniePods {
			fmt.Println(erniePod.(*v1.Pod).Name)
		}
	*/

	pods, err := index.ByIndex("namespace", "default")
	if err != nil {
		panic(err)
	}

	for _, pod := range pods {
		fmt.Println(pod.(*v1.Pod).Name)
	}
}
