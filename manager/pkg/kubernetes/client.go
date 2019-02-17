package kubernetes

import (
	root "github.com/c-o-l-o-r/watchtower/manager/pkg"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
)

type Client struct {
	deployments appsv1.DeploymentInterface
}

func NewClient(config *root.KubernetesConfig) (*Client, error) {
	clusterConfig, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		panic(err.Error())
	}

	return &Client{deployments: clientset.AppsV1().Deployments(config.Namespace)}, err
}
