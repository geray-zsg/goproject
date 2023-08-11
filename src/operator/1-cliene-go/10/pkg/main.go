package main

import (
	"fmt"
	"goproject/src/operator/1-cliene-go/10/pkg/informer"

	"k8s.io/apimachinery/pkg/labels"
)

func main() {

	stopCh := make(chan struct{})

	err := informer.NewSharedInformerFactory(stopCh)
	if err != nil {
		panic(err)
	}

	items, err := informer.Get().Core().V1().Pods().Lister().List(labels.Everything())
	if err != nil {
		panic(err)
	}
	for _, v := range items {
		fmt.Printf("namespace: %v,name: %v", v.Namespace, v.Name)
	}
}
