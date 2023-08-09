/**
 * @Author: Geray
 * @Date: 2023/8/9 20:23:17
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/9 20:23:17
 * Description:
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
	// 1、加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../1-restClient/kubeconfig")
	if err != nil {
		panic(err)
	}

	// 2、实例化客户端对象
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 3、配置我们需要调用的GVR
	gvr := schema.GroupVersionResource{
		Group:    "", //不需要写，因为无名资源组，也就是 core 资源组
		Version:  "v1",
		Resource: "pods",
	}

	// 4、发送请求，且得到返回结果（非结构化）
	unStrucData, err := dynamicClient.Resource(gvr).Namespace("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// 5、转换为结构化
	// Resource，基于 gvr生成了一个针对于资源的客户端，也可以称之为动态资源客户端 dynamicResourceClient
	// namespace，指定一个可操作的命名空间。同时他是dynamicResourceClient的方法
	// List，首先是通过RESTClient 调用K8s APIServer 的接口返回了Pod 数据，返回的数据格式 是二进制的Jdon格式，然后通过一些列解析方法，转换成 unstructured.UnstructuredList。
	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(
		unStrucData.UnstructuredContent(),
		podList,
	)
	if err != nil {
		panic(err)
	}

	// 6、遍历
	fmt.Println(podList)
	for _, item := range podList.Items {
		fmt.Printf("namespace: %v, name: %v \n", item.Namespace, item.Name)
	}
}
