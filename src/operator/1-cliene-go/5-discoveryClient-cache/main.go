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
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery/cached/disk"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1、加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../1-restClient/kubeconfig")
	if err != nil {
		panic(err)
	}

	// 2、实例化客户端对象，本地客户端负责将GVR数据，缓存到本地文件夹（kubectl也就是使用的这种方式将缓存放在.kube/cache中）
	cacheDiscoveryClient, err := disk.NewCachedDiscoveryClientForConfig(config, "./cache/discovery", "./cache/http", time.Minute*60)
	if err != nil {
		panic(err)
	}

	_, apiResources, err := cacheDiscoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err.Error())
	}

	// 1.先从缓存文件中找GVR数据，有则返回没有则调用 APIServer
	// 2.调用APIServer 获取GVR数据
	// 3.将获取到的 GVR 数据缓存到本地，然后返回给客户端

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
