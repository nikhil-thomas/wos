package main

import (
    "context"
    "flag"
    "fmt"
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    restclient "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
)

func main() {
    config := getKubeconfigOrPanic()
    clientSet, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err)
    }
    podClient := clientSet.CoreV1().Pods("")
    listOptions := metav1.ListOptions{
        Limit: 50,
    }
    for {
        pods, err := podClient.List(context.Background(), listOptions)
        if err != nil {
            panic(err.Error())
        }
        printPodList(pods)
        if pods.Continue == "" {
            fmt.Println("continue is \"\"")
            break
        }
        listOptions.Continue = pods.Continue
        fmt.Println("--------------------------------------------")
    }
    pods, err := podClient.List(context.Background(), listOptions)

    fmt.Printf("continue: %s", pods.Continue)
    if err != nil {
        panic(err.Error())
    }
}

func printPodList(pods *corev1.PodList) {
    fmt.Printf("\nPODS\n\n")
    for i, pod := range pods.Items {
        fmt.Printf("%d. %s/%s\n", i+1, pod.GetNamespace(), pod.GetName())
    }
}

func getKubeconfigOrPanic() *restclient.Config {

    kubeconfigPath := flag.String("kubeconfig", "", "path to kubeconfig file")
    flag.Parse()
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigPath)
    if err != nil {
        panic(err.Error())
    }
    return config
}
