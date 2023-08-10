/**
* @Author: Geray
* @Date: 2023/8/10 19:27:07
* @LastEditors: Geray
* @LastEditTime: 2023/8/10 19:27:07
* Description:

IndexFunc∶索引器函数，用于计算一个资源对象的索引值列表，可以根据需求定义其他的，比如根据Label标签、Annotation等属性来生成索引值列表。
Index:存储数据，要查找某个命名空间下面的 Pod，那就要让Pod按照其命名空间进行索引，对应的Index类型就是 map[nomespce]sets.pod。
Indexers︰存储索引器，key 为索引器名称， value为索引器的实现函数，例如: mapl["namespoace" ]HetolNamespoceIndexFunc,
Indices∶存储缓存器，key 为索引器名称，value 为缓存的数据，例如: map[ "namespace "]map [namespace]sets.pod。

* Copyright: Copyright (©)}) 2023 Geray. All rights reserved.
*/
package main

import (
	"fmt"

	v1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

/*
	1.实现两个索引器函数，分别基于Namespace、NodeName，资源对象Pod
*/

// namespace索引器
func NamespaceIndexFunc(obj interface{}) (result []string, err error) {

	pod, ok := obj.(*v1.Pod) // 指针类型
	if !ok {
		return nil, fmt.Errorf("类型错误 %v", err)
	}
	// 索引值列表
	result = []string{pod.Namespace}
	return
}

// NodeName索引器
func NodeNameIndexerFunc(obj interface{}) (result []string, err error) {

	pod, ok := obj.(*v1.Pod)
	if !ok {
		return nil, fmt.Errorf("类型错误 %v", err)
	}

	// 索引值列表
	result = []string{pod.Spec.NodeName}
	return

}

func main() {
	// 实例一个indexer对象
	index := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{
		"namespace": NamespaceIndexFunc,
		"nodeName":  NodeNameIndexerFunc,
	})

	// 模拟数据
	pod1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "index-pod-1",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			NodeName: "node1",
		},
	}
	pod2 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "index-pod-2",
			Namespace: "kube-system",
		},
		Spec: v1.PodSpec{
			NodeName: "node2",
		},
	}
	pod3 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "index-pod-3",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			NodeName: "node3",
		},
	}

	// 加入数据到Indexer中

	_ = index.Add(pod1)
	_ = index.Add(pod2)
	_ = index.Add(pod3)

	// 通过索引器函数查询一下数据：pod
	pods, err := index.ByIndex("namespace", "default")
	if err != nil {
		panic(err)
	}

	for _, pod := range pods {
		// 断言
		fmt.Println(pod.(*v1.Pod).Name)

	}

	fmt.Println("----------------------------------")
	// 通过索引器函数查询一下数据：NodeName
	pods, err = index.ByIndex("nodeName", "node2")
	if err != nil {
		panic(err)
	}

	for _, pod := range pods {
		// 断言
		fmt.Println(pod.(*v1.Pod).Name)

	}

	/*

	 */

}
