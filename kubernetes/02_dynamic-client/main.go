package main

import (
	"bufio"
	"flag"
	"fmt"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/apimachinery/pkg/runtime/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/util/retry"
	"os"
	"path/filepath"
	"k8s.io/client-go/util/homedir"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to kubeconfig file")
	}

	flag.Parse()
	namespace := "default"

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentRes := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}

	deployment := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind": "Deployment",
			"metadata": map[string]interface{}{
				"name": "demo-deployment",
				//"labels": map[string]interface{}{
				//	"app": "deploy-test",
				//	"label2": "2134",
				//},
			},
			"spec": map[string]interface{}{
				"replicas": 2,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "demo",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": "demo",
						},
					},
					"spec": map[string]interface{}{
						"containers": []map[string]interface{}{
							{
								"name": "web",
								"image": "nginx:1.12",
								"ports": []map[string]interface{}{
									{
										"name": "http",
										"protocol": "TCP",
										"containerPort": 80,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("creating deployment")
	result, err := client.Resource(deploymentRes).Namespace(namespace).Create(deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q. \n", result.GetName())

	prompt()
	fmt.Println("Updating deployment...")

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := client.Resource(deploymentRes).Namespace(namespace).Get("demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}

		if err := unstructured.SetNestedField(result.Object, int64(1), "spec", "replicas"); err != nil {
			panic(fmt.Errorf("failed to set replica value: %v", err))
		}

		//extract spec containers
		containers, found, err := unstructured.NestedSlice(result.Object, "spec", "template", "spec", "containers")
		if err != nil || !found || containers == nil{
			panic(fmt.Errorf("deployment containers not found or error in spec: %v", err))
		}
		//update containers[0] image
		if err := unstructured.SetNestedField(containers[0].(map[string]interface{}), "nginx:1.13", "image"); err != nil {
			panic(err)
		}
		if err := unstructured.SetNestedField(result.Object, containers, "spec", "template", "spec", "containers"); err != nil {
			panic(err)
		}
		_,updateErr := client.Resource(deploymentRes).Namespace(namespace).Update(result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("update failed: %v", retryErr))
	}
	fmt.Println("updated deployment")
	prompt()
	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := client.Resource(deploymentRes).Namespace(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _,d :=range list.Items {
		replicas, found, err := unstructured.NestedInt64(d.Object, "spec", "replicas")
		if err != nil || !found {
			fmt.Printf("Replicas not found for deployment %s, err %s", d.GetName(), err)
			continue
		}
		fmt.Printf(" * %s (%d replicas)\n", d.GetName(), replicas)
	}
	prompt()

	fmt.Println("Labels")

	retryErr = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := client.Resource(deploymentRes).Namespace(namespace).Get("demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}

		labels, found, err := unstructured.NestedMap(result.Object, "metadata", "labels")
		if err != nil{
			panic(fmt.Errorf("deployment labels find error in spec: %v", err))
		}
		fmt.Println("Found", found)
		fmt.Printf("%+v\n", labels)
		if labels == nil {
			labels = map[string]interface{}{}
		}
		labels["new-label"] = "true"
		err = unstructured.SetNestedMap(result.Object, labels, "metadata", "labels")
		if err != nil{
			panic(fmt.Errorf("deployment labels set error in spec: %v", err))
		}
		labels, found, err = unstructured.NestedMap(result.Object, "metadata", "labels")
		if err != nil{
			panic(fmt.Errorf("deployment labels find error in spec: %v", err))
		}
		fmt.Println("Found", found)
		fmt.Printf("%+v\n", labels)
		//update containers[0] image
		//if err := unstructured.SetNestedField(containers[0].(map[string]interface{}), "nginx:1.13", "image"); err != nil {
		//	panic(err)
		//}
		//if err := unstructured.SetNestedField(result.Object, containers, "spec", "template", "spec", "containers"); err != nil {
		//	panic(err)
		//}
		//_,updateErr := client.Resource(deploymentRes).Namespace(namespace).Update(result, metav1.UpdateOptions{})
		return nil
	})
	if retryErr != nil {
		panic(fmt.Errorf("update failed: %v", retryErr))
	}

	prompt()

	retryErr = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := client.Resource(deploymentRes).Namespace(namespace).Get("demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}

		copyDeployment := result.DeepCopy()
		copyDeployment.SetName("demo-deployment-copy")
		unstructured.RemoveNestedField(copyDeployment.Object, "metadata", "resourceVersion")

		_, err = client.Resource(deploymentRes).Namespace(namespace).Create(copyDeployment, metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}
		return nil
	})
	if retryErr != nil {
		panic(fmt.Errorf("update failed: %v", retryErr))
	}

	prompt()

	fmt.Println("Deleting Deployment")
	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := &metav1.DeleteOptions{
		PropagationPolicy:  &deletePolicy,
	}
	if err := client.Resource(deploymentRes).Namespace(namespace).Delete("demo-deployment", deleteOptions); err != nil {
		panic(err)
	}
	if err := client.Resource(deploymentRes).Namespace(namespace).Delete("demo-deployment-copy", deleteOptions); err != nil {
		panic(err)
	}
	fmt.Println("Deleted deployment")
}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}
