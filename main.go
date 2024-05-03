package main

import (
	"context"
	"flag"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	crdclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	_ "k8s.io/code-generator"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	home := homedir.HomeDir()
	if home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	crdClient, err := crdclientset.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	customCRD := v1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "Shaikat.crd.com",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "crd.com",
			Versions: []v1.CustomResourceDefinitionVersion{
				{
					Name:    "v1alpha1",
					Served:  true,
					Storage: true,
					Schema: &v1.CustomResourceValidation{
						OpenAPIV3Schema: &v1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1.JSONSchemaProps{
								"spec": {
									Type: "object",
									Properties: map[string]v1.JSONSchemaProps{
										"name": {
											Type: "string",
										},
										"replicas": {
											Type: "integer",
										},
										"container": {
											Type: "object",
											Properties: map[string]v1.JSONSchemaProps{
												"image": {
													Type: "string",
												},
												"port": {
													Type: "integer",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			Scope: "Namespaced",
			Names: v1.CustomResourceDefinitionNames{
				Kind:     "Shaikat",
				Plural:   "shaikatt",
				Singular: "shaikat",
				ShortNames: []string{
					"skt",
					"shkt",
				},
				Categories: []string{
					"all",
				},
			},
		},
	}
	ctx := context.TODO()

}
