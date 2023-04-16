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

	flag.Var(&kubeconfig, "kubeconfig", "specify the path to the kubeconfig")

	flag.Parse()

	if kubeconfig.String() == "" {
		if err := kubeconfig.Set(kubevar.Default()); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	fmt.Printf("Config:%v", kubeconfig.String())

	// get the cluster-info
	svcs, err := kubeconfig.Clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, svc := range svcs.Items {
		fmt.Println(svc.ObjectMeta.Name)
	}
	return
}
