package controller

import "k8s.io/client-go/kubernetes"

type Controller interface {
	kubeclientset  kubernetes.Interface
	sampleclientset clientset.Interface
	deploy

}
