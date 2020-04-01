package main

// k8s.io对应https://github.com/kubernetes/*
import (
	"flag"
	"fmt"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/transport/spdy"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	//restclient "k8s.io/client-go/rest"
)

func main() {
	clientSet := getClientSet()

	//getAllPods(clientSet)

	//getOnePod(clientSet)

	//getAllDeployment(clientSet)

	//getOneDeployment(clientSet)

	//getPodLog(clientSet)

	//execPod(clientSet)
	test(clientSet)
}

func test(clientset *kubernetes.Clientset){
	bytes, err := clientset.
		RESTClient().
		Get().
		RequestURI("https://192.168.40.160:6443/apis/metrics.k8s.io/v1beta1/namespaces/istio-system/pods/istio-telemetry-6587975b4c-v8m8r").
		DoRaw()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func execPod(clientset *kubernetes.Clientset){
	//remotecommand.NewSPDYExecutor()
	//client := clientset.RESTClient()
	//closer, err := client.Post().RequestURI("https://192.168.40.121:6443/api/v1/namespaces/logging/pods/fluentd-es-ssxds/exec?command=bash&container=fluentd-es&stdin=true&stdout=true&tty=true").Stream()
	//if err != nil {
	//	panic(err)
	//}
	//defer closer.Close()
	//io.Copy(os.Stdout,closer)

	var tsq remotecommand.TerminalSizeQueue
	so := remotecommand.StreamOptions{
		Stdout:os.Stdout,
		Stdin:os.Stdin,
		Stderr:os.Stdout,
		Tty:true,
		TerminalSizeQueue:tsq,
	}
	var transport http.RoundTripper
	var upgrader spdy.Upgrader
	var url  = url.URL{
		Scheme:"https",
		Host:"192.168.40.160:6443",
		///api/v1/namespaces/istio/pods/cloudapiserver-86b8d5768f-dc5cr/exec?command=bash&container=cloudapiserver&stdin=true&stdout=true&tty=true
		Path:"/api/v1/namespaces/logging/pods/fluentd-es-ssxds/exec?container=fluentd-es",
	}
	executor, err := remotecommand.NewSPDYExecutorForProtocols(transport, upgrader, "POST", &url, "http")
	if err != nil {
		panic(err)
	}
	err = executor.Stream(so)
	if err != nil {
		panic(err)
	}
}

func getPodLog(clientSet *kubernetes.Clientset) {
	var i int64
	i = 100
	opts := &v1.PodLogOptions{
		Container: "fluentd-es",
		Follow:true,
		TailLines: &i,
	}

	request := clientSet.CoreV1().Pods("logging").GetLogs("fluentd-es-ssxds", opts)

	// 追踪日志　kubectl logs -f
	closer, err := request.Stream()
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	io.Copy(os.Stdout,closer)

	// 静态日志　kubectl logs
	//bytes, err := request.Do().Raw()
	//if err != nil {
	//	panic(err)
	//}
	//	print(string(bytes))

}

func getOneDeployment(clientSet *kubernetes.Clientset) {
	deployment, err := clientSet.AppsV1().Deployments("istio").Get("cloudapiserver", metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Println("deployment not found")
	}
	if errors.IsAlreadyExists(err) {
		fmt.Println("deployment已经存在")
	}
	fmt.Println(deployment)
}

func getAllDeployment(clientSet *kubernetes.Clientset) {
	deploys, err := clientSet.AppsV1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for i, v := range deploys.Items {
		fmt.Printf("id=%d \t nameSpace:%s \t name:%s \n", i, v.Namespace, v.Name)
	}
}

func getOnePod(clientSet *kubernetes.Clientset) {
	namespace := "kube-system"
	pod := "calico-node-8cc5n"
	_, err := clientSet.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	}
}

func getAllPods(clientSet *kubernetes.Clientset) {
	pods, err := clientSet.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for i, v := range pods.Items {
		fmt.Printf("id: %d \t NameSpace: %s \t Name: %s\n", i, v.Namespace, v.Name)
	}
}

func getClientSet() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientSet
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
