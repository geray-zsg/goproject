/**
 * @Author: Geray
 * @Date: 2023/8/11 12:43:45
 * @LastEditors: Geray
 * @LastEditTime: 2023/8/11 12:43:45
 * Description: Client-go调用封装
 * Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
 */
package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

/*
	Gin + Client-go 开发示例
	封装client-go

*/
// 定义结构体
type Clients struct {
	clienetSet kubernetes.Interface
}

// 客户端
func NewClients() (clients Clients) {

	// 加载配置生成配置对象
	config, err := clientcmd.BuildConfigFromFlags("", "kubeconfig")
	if err != nil {
		return
	}

	// 实例化客户端
	clients.clienetSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		return
	}

	return
}

// 定义方法：由于上面结构体重的clientSet是内部方法，不能对外提供，这里使用方法对外暴露
func (c *Clients) ClientSet() kubernetes.Interface {
	return c.clienetSet
}
