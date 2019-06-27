package main

// k8s.io对应https://github.com/kubernetes/*
import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	clientSet := getClientSet()

	//getAllPods(clientSet)

	//getOnePod(clientSet)

	//getAllDeployment(clientSet)

	getOneDeployment(clientSet)
}

func getOneDeployment(clientSet *kubernetes.Clientset) {
	deployment, err := clientSet.AppsV1().Deployments("kube-system").Get("cloudapiservera", metav1.GetOptions{})
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
