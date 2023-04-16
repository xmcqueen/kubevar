package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"q/cl/kubevar"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	kubeconfig := kubevar.Kubeconfig{}

	labelSelector := flag.String("l", "kubernetes.io/cluster-service=true", "labelSelector for the ListOptions for the API call")
	flag.Var(&kubeconfig, "kubeconfig", "specify the path to the kubeconfig")

	flag.Parse()

	if kubeconfig.String() == "" {
		if err := kubeconfig.Set(kubevar.Default()); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	fmt.Printf("Kubeconfig: %v\n", kubeconfig.String())

	// get the cluster-info
	svcs, err := kubeconfig.Clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{LabelSelector: *labelSelector})
	if err != nil {
		panic(err)
	}

	for _, svc := range svcs.Items {
		fmt.Println(svc.ObjectMeta.Name)
	}
	return
}
