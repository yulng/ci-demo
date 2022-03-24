package ip

import (
	"ci-demo/test/e2e/ging"
	"context"
	"fmt"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

var _ = ginkgo.Describe("[Test IP Allocation]", func() {
	g := ging.NewGinG("[Test IP Allocation]", fmt.Sprintf("%v/.kube/config", homedir.HomeDir()))

	ginkgo.Describe("test describe", func() {
		ginkgo.It("test It", func() {
			namespaces, err := g.KubeClientSet.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
			if err != nil {
				klog.Fatalf("[error when get namespaces]", err)
				return
			}
			for _, ns := range namespaces.Items {
				fmt.Printf("%v", ns)
			}

			dsc := g.GetDescribe()
			gomega.Expect(dsc).Should(gomega.Equal("test-It"))
		})
	})
})
