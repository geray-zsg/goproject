/**
 * @Author: Geray
 * @Date: 2023/8/9 20:23:17
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/9 20:23:17
 * Description:	获取资源组
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1、加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../1-restClient/kubeconfig")
	if err != nil {
		panic(err)
	}

	// 2、实例化客户端对象
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}

	// 3、发送请求，获取GVR数据
	_, apiResources, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}

	// 6、遍历，解析GV数据
	// fmt.Println(apiResources)
	for _, list := range apiResources {
		gv, err := schema.ParseGroupVersion(list.APIVersion)
		if err != nil {
			panic(err)
		}
		for _, resource := range list.APIResources {
			fmt.Printf("name: %v, group:%v, version: %v} \n", resource.Name, gv.Group, gv.Version)
		}
	}
}
