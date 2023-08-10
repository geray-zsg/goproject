/**
 * @Author: Geray
 * @Date: 2023/8/9 20:23:17
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/9 20:23:17
 * Description:	List-watch（List之前的都是获取到全量数据，这里演示watch）
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
	// 1、加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../1-restClient/kubeconfig")
	if err != nil {
		panic(err)
	}

	// 2、实例化客户端对象，clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 3、调用监听方法
	w, err := clientSet.AppsV1().Deployments("default").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// 死循环进行持续监听
	fmt.Println("watch开启...")
	for {
		select {
		case e, _ := <-w.ResultChan():
			fmt.Println(e.Type, e.Object)
			// e.Type：表示事件变化的类型，Added，DELETE
			// e.Object：表示变化后的数据，可以启动后尝试操作k8s资源会有相关信息输出（比较大）
		}
	}

}
