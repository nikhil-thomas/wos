package main

import (
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kcPath := filepath.Join("/", "home", "ncoder", ".kube", "config")
	kubeconfig := flag.String("kubeconig", kcPath, "kubeconfig path")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	pods, err := clientSet.CoreV1().Pods("kube-system").List(metav1.ListOptions{})
	for i, pod := range pods.Items {
		fmt.Println("pod", i, ":", pod.Name, pod.ObjectMeta.GetResourceVersion(), pod.GroupVersionKind().Kind)
	}

	deps, err := clientSet.AppsV1().Deployments("kube-system").List(metav1.ListOptions{})
	for i, dep := range deps.Items {
		fmt.Println("dep", i, ":", dep.Name, dep.ObjectMeta.GetResourceVersion(), dep.GroupVersionKind().Kind)
	}
}
