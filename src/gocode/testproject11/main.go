/**
 * @Author: Geray
 * @Date: 2023/7/22 11:02:25
 * @LastEditors: Geray
 * @LastEditTime: 2023/7/22 11:02:25
 * Description: 获取 kube-system 这个名称空间下的 Pod 列表
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	/*
		1. k8s的配置文件kubeconfig
		2. 保证开发及能通过这个配置文件连接到 k8s 集群

	*/

	// 加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "kubeconfig")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 2. 配置API路径
	config.APIPath = "api" // pods , /api/v1/pods
	// config.APIPath = "apis" // deployments , /apis/v1/namespaces/{namespace}/deployments/{deployment}

	// 3. 配置分组版本
	config.GroupVersion = &corev1.SchemeGroupVersion // 无名资源组，group: "", version: "v1"

	// 4. 配置数据端的编码工具
	config.NegotiatedSerializer = scheme.Codecs

	// 5. 实例化 RESTClient 对象
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}

	// 6. 定义接收返回值的变量
	result := &corev1.PodList{}

	// 7. 跟APIServer 交互
	/*
		Get: 定义了请求方式，返回了一个 Request 结构体对象。这个 Request 结构体对象，就是构建访问 APIServer 请求用的。
		依次执行了 Namespace，Resource，VersionedParams，构建与 APIServer 交互的参数
		Do方法通过 request 发起请求，然后通过 transformResponse 解析请求返回，并绑定到对应资源对象的结构体对象上，这里的话，就表示是corev1,.PodList的对象
		request 先是检查了有没有可用的 Client，在这里开始调用 net/http 包的功能
	*/
	restClient.
		Get().                                                         // Get请求方式
		Namespace("kube-system").                                      // 指定 名称空间
		Resource("pods").                                              // 指定需要查询的资源， 传递资源名称
		VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec). //参数及参数的序列化工具
		Do(context.TODO()).                                            // 触发请求
		Into(result)                                                   //写入返回结果

	// 输出结果（PodList结构体，阅读性很差）
	// fmt.Println(result)

	// 格式化输出结果
	for _, item := range result.Items {
		// \n Printf 没有换行
		fmt.Printf("namespace: %v, name: %v \n", item.Namespace, item.Name)
	}
}
