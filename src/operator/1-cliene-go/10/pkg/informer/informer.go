package informer

import (
	"goproject/src/operator/1-cliene-go/10/pkg/client"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
)

var sharedInformerFactory informers.SharedInformerFactory

func NewSharedInformerFactory(stopCh <-chan struct{}) (err error) {

	var (
		clients client.Clients
	)

	// 1、加载客户端
	clients = client.NewClients()

	// 2、实例化sharedInformerFactory
	sharedInformerFactory = informers.NewSharedInformerFactoryWithOptions(clients.ClientSet(), time.Second*60)

	// 3、启动informer
	gvrs := []schema.GroupVersionResource{
		{Group: "", Version: "v1", Resource: "pods"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "namespaces"},

		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "apps", Version: "v1", Resource: "statefulsets"},
		{Group: "apps", Version: "v1", Resource: "daemonsets"},
	}

	for _, v := range gvrs {
		// 创建informer
		_, err = sharedInformerFactory.ForResource(v)
		if err != nil {
			return
		}
	}

	// 启动所有创建的 informer
	sharedInformerFactory.Start(stopCh)

	// 等待所有的informer全量同步数据完成
	sharedInformerFactory.WaitForCacheSync(stopCh)

	return

}

func Get() informers.SharedInformerFactory {
	return sharedInformerFactory
}
