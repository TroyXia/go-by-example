package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
)

func main() {
	// generate RESTClient config
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hxia/.kube/config")
	if err != nil {
		panic(err)
	}

	// add the required information for config
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// generate RESTClient
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get pod list via RESTClient
	// request:
	// 	Scheme = "https"
	//  Host = "127.0.0.1:54223"
	//  Path = "/api/v1/namespaces/default/pods"
	//  RawQuery = "limit=500"
	pods := &corev1.PodList{}
	if err := restClient.Get().
		Namespace("default").Resource("pods").VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(pods); err != nil {
		panic(err)
	}

	for _, d := range pods.Items {
		fmt.Printf("NAMESPACE: %v \t NAME: %v \t STATUS:%+v\n", d.Namespace, d.Name, d.Status.Phase)
	}
}
