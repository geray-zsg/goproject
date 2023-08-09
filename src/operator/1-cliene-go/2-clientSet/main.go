package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1、加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../1-restClient/kubeconfig")
	if err != nil {
		panic(err)
	}

	// 2、实例化RESTClient对象
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 3、获取结果
	podList, err := clientSet.
		CoreV1().                                  // 返回CoreV1Client 实例
		Pods("kube-system").                       // 指定查询的资源以及指定资源的namespace，namespace如果为空，则表示查询的所有namespace
		List(context.TODO(), metav1.ListOptions{}) // 在这里表示查询Pod列表

	if err != nil {
		panic(err)
	}

	// 4、遍历列表获取值
	for _, item := range podList.Items {
		fmt.Printf("namespace: %v, podName: %v \n", item.Namespace, item.Name)
	}
}
