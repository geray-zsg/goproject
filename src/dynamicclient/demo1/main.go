/**
 * @Author: Geray
 * @Date: 2023/7/22 12:22:20
 * @LastEditors: Geray
 * @LastEditTime: 2023/7/22 12:22:20
 * Description: 使用dynamicClient获取 pod 列表（dynamicClient可以获取到k8s中所有资源，包括自定义的）
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1. 加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../../clientset/demo1/kubeconfig")
	if err != nil {
		panic(err.Error())
	}

	// 2. 实例化客户端对象，这里是实例化 动态客户端对象
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3. 配置我们需要调用的 GVR
	gvr := schema.GroupVersionResource{
		Group:    "", // 不需要写，因为是无名的资源组，也就是 core 资源组
		Version:  "v1",
		Resource: "pods",
	}

	// 4. 发起请求，并得到返回的结果（非结构化数据）
	/*
		Resource基于GVR 生成了一个针对域资源的客户端，也可以称之为动态资源客户端，dynamicResourceClient
		Namespace：指定了一个可操作性的名称空间，同时他是 dynamicResourceClient 的方法
		List：首先是通过RESTClient 调用 k8s APIServer 的接口返回 Pod的数据，返回的数据格式是 二进制的 json格式
			然后通过一系列 的解析方法，转换成 unstructured.UnstructuredList

	*/
	unStructData, err := dynamicClient.Resource(gvr).Namespace("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(unStructData)

	// 5. unStructData 转换为结构化数据
	// config.GroupVersion = &corev1.SchemeGroupVersion
	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(
		unStructData.UnstructuredContent(),
		podList,
	)
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range podList.Items {
		fmt.Printf("namespace: %v, name: %v \n", pod.Namespace, pod.Name)
	}

}
