package main

import (
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"log"
	"path/filepath"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err)
	} // 初始化 client

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err.Error())
	}
	stopper := make(chan struct{})

	defer close(stopper) // 初始化 informer

	factory := informers.NewSharedInformerFactory(clientset, 0)
	nodeInformer := factory.Core().V1().Nodes()
	namespaceInformer := factory.Core().V1().Namespaces()
	informer := nodeInformer.Informer()
	nsInformer := namespaceInformer.Informer()
	defer runtime.HandleCrash() // 启动 informer，list & watch

	go factory.Start(stopper) // 从 apiserver 同步资源，即 list

	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	} // 使用自定义 handler

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: func(interface{}, interface{}) { fmt.Println("update not implemented") }, // 此处省略 workqueue 的使用
		DeleteFunc: func(interface{}) { fmt.Println("delete not implemented") },
	}) // 创建 lister

	nsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("add ns")
			namespace := obj.(*corev1.Namespace)
			println(namespace.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {

		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete ns")
			namespace := obj.(*corev1.Namespace)
			println(namespace.Name)
		},
	})
	nodeLister := nodeInformer.Lister()
	// 从 lister 中获取所有 items
	nodeList, err := nodeLister.List(labels.Everything())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("nodelist:", nodeList)
	<-stopper
}
func onAdd(obj interface{}) {

	node := obj.(*corev1.Node)
	fmt.Println("add a node:", node.Name)
}
