package ip

import (
	"ci-demo/test/e2e/ging"
	"context"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

var _ = Describe("[Test IP Allocation]", func() {
	g := ging.NewGinG("[Test IP Allocation]", fmt.Sprintf("%v/.kube/config", os.Getenv("HOME")))

	Describe("test describe", func() {
		It("test It", func() {
			namespaces, err := g.KubeClientSet.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
			if err != nil {
				klog.Fatalf("[error when get namespaces]", err)
				return
			}
			for _, ns := range namespaces.Items {
				fmt.Printf("%v", ns)
				fmt.Printf("happy")
			}

			dsc := g.GetDescribe()
			Expect(dsc).Should(Equal("test-It"))
		})
	})
})
