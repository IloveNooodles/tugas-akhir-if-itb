package main

import (
	"context"

	"github.com/ilovenooodles/poc-remote-deployment/controller"
	"github.com/sirupsen/logrus"
	// appsv1 "k8s.io/api/apps/v1"
	// apiv1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/util/retry"
)

func main() {
	l := logrus.New()
	ctx := context.Background()
	kc, err := controller.New(l)
	if err != nil {
		panic(err)
	}

	kc.SwitchCluster("prod-cluster-example")

	// // Labeling all nodes
	// nodeInterface := kc.GetNodeInterface()
	// nodes, err := nodeInterface.List(context.TODO(), v1.ListOptions{})
	// if err != nil {
	// 	panic(err)
	// }

	err = kc.LabelNodes(ctx, "test-cluster-1-worker", "bismillah", "sukses")
	if err != nil {
		kc.Logger.Errorf("error when labelling nodes err: %s", err)
		panic(err)
	}

	// for idx, node := range nodes.Items {
	// 	fmt.Printf("[%d] Current Node: %s\n", idx, node.Name)
	// key := fmt.Sprintf("node-%d", idx)
	// val := fmt.Sprintf("node-worker-value-%d", idx)
	// err := kc.LabelNodes(ctx, node.Name, key, val)
	// if err != nil {
	// 	kc.Logger.Errorf("error when labelling nodes err: %s", err)
	// 	panic(err)
	// }
	// }

	// labels := make(map[string]string)
	// labels["app"] = "demo"

	// target := make(map[string]string)
	// target["contoh"] = "inideployment"

	// params := controller.DeployParams{
	// 	Replica: 3,
	// 	Name:    "web2",
	// 	Image:   "nginx:latest",
	// 	Labels:  labels,
	// 	Targets: target,
	// }

	// res, err := kc.Deploy(ctx, params)
	// if err != nil {
	// 	fmt.Println("SUDAH ADA KAH")
	// }

	// logDetail(res)

	// err = kc.Delete(ctx, params)
	// logDetail(err)
}

func logDetail(v any) {
	// log.Printf("%#v\n", v)
}
