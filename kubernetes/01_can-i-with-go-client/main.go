package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"

	authorizationv1 "k8s.io/api/authorization/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfigpath := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)
	if err != nil {
		panic(err.Error())
	}

	//create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	authtest := &authorizationv1.SubjectAccessReview{
		Spec: authorizationv1.SubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:      "create",
				Resource:  "pods",
				Namespace: "ns-a",
			},
			User: "system:serviceaccount:ns-a:sa1",
		},
	}

	authtest, err = clientset.AuthorizationV1().SubjectAccessReviews().Create(authtest)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(authtest.Status.Allowed)
}
