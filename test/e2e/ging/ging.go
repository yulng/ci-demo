package ging

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"strings"
)

type GinG struct {
	Describe      string
	KubeClientSet kubernetes.Interface
	KubeConfig    *rest.Config
}

func NewGinG(describe, kubeConfigPath string) GinG {
	// generate instance
	g := GinG{Describe: describe}
	// client-go
	conf, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		klog.Fatalf("[error when get kubeConfigPath]%v", err)
	}
	g.KubeConfig = conf

	kubeClient, err := kubernetes.NewForConfig(conf)
	if err != nil {
		klog.Fatalf("[error when new kubeClientSet]%v", err)
	}
	g.KubeClientSet = kubeClient
	return g
}

func (g *GinG) GetDescribe() string {
	return strings.Replace(CurrentGinkgoTestDescription().TestText, " ", "-", -1)
}
