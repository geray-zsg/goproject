/**
 * @Author: Geray
 * @Date: 2023/7/22 12:03:58
 * @LastEditors: Geray
 * @LastEditTime: 2023/7/22 12:03:58
 * Description:	使用了clientset封装，相比demo1 代码更加简介，简单，易读，易写
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1. 加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../demo1/kubeconfig")
	if err != nil {
		panic(err.Error())
	}

	//2. 实例化 clientset 对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	/*
		   CoreV1 返回 CoreV1Client 实例对象
		   Pods 调用了newPods 函数，该函数返回的是 PodInterface 对象，
		   	PodInterface 对象 实现了 Pods 资源相关的全部方法，
			同时newPods 里面还将 RESETClient实例对象复制给了对应的CLient属性
		   List 内使用了 RestClient 与 k8s APIServer 进行了交互
	*/
	pods, err := clientset.
		CoreV1().            //返回CoreV1Client 实例
		Pods("kube-system"). //查询命名空间 pod 列表
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(pods)

	for _, pod := range pods.Items {
		fmt.Printf("namespace: %v, name: %v \n", pod.Namespace, pod.Name)
	}
}
