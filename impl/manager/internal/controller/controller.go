package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	apiappsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	kubeerror "k8s.io/apimachinery/pkg/api/errors"
	apimetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/retry"
)

var ()

type KubernetesController struct {
	Logger    *logrus.Logger
	Config    clientcmd.ClientConfig
	ClientSet *kubernetes.Clientset
}

// Create new kubernetes Controller
func New(l *logrus.Logger) (*KubernetesController, error) {
	cfg, err := generateKubeConfig()
	if err != nil {
		err := fmt.Errorf("kube: error when config path err: %s", err)
		l.Error(err)
		return nil, err
	}

	rawCfg, err := cfg.RawConfig()
	if err != nil {
		err := fmt.Errorf("kube: error when reading raw cfg err: %s", err)
		l.Error(err)
		return nil, err
	}

	restCfg, err := cfg.ClientConfig()
	if err != nil {
		err := fmt.Errorf("kube: error when reading rest cfg err: %s", err)
		l.Error(err)
		return nil, err
	}

	clienset, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		err := fmt.Errorf("kube: error when creating clientset err: %s", err)
		l.Error(err)
		return nil, err
	}

	l.Infof("kube: success connected to context: %s", rawCfg.CurrentContext)
	return &KubernetesController{
		Logger:    l,
		Config:    cfg,
		ClientSet: clienset,
	}, nil
}

// Return config instance
func (k *KubernetesController) GetConfig() clientcmd.ClientConfig {
	return k.Config
}

// Return raw config instance
func (k *KubernetesController) GetRawConfig() (api.Config, error) {
	rawCfg, err := k.Config.RawConfig()
	if err != nil {
		k.Logger.Errorf("kube: error when getting raw config err: %s", err)
		return rawCfg, err
	}
	return rawCfg, nil
}

// Return rest config instnance
func (k *KubernetesController) GetRestConfig() (*rest.Config, error) {
	restCfg, err := k.Config.ClientConfig()
	if err != nil {
		k.Logger.Errorf("kube: error when getting rest config err: %s", err)
		return restCfg, err
	}

	return restCfg, nil
}

// Get Nodes from CoreV1
func (k *KubernetesController) GetNodeInterface() v1.NodeInterface {
	return k.ClientSet.CoreV1().Nodes()
}

// Will switch kubernetes context with ctx.
//
// Will iterate through all of available context and validate
// if the ctx is exists in the list. If found, will change
// current context to ctx.
func (k *KubernetesController) SwitchContext(ctx string) error {
	configAccess := k.Config.ConfigAccess()
	rawConfig, err := k.Config.RawConfig()
	if err != nil {
		err := fmt.Errorf("error %s, getting starting config", err.Error())
		k.Logger.Error(err)
		return err
	}

	currentContext := rawConfig.CurrentContext
	k.Logger.Infof("current context %s", currentContext)

	if currentContext == ctx {
		return nil
	}

	err = validateContext(&rawConfig, ctx)
	if err != nil {
		err := fmt.Errorf("error %s, validating kubeconfig", err.Error())
		k.Logger.Error(err)
		return err
	}

	rawConfig.CurrentContext = ctx
	err = clientcmd.ModifyConfig(configAccess, rawConfig, true)
	if err != nil {
		err := fmt.Errorf("error %s, modifying config", err.Error())
		k.Logger.Error(err)
		return err
	}

	cfg := clientcmd.NewNonInteractiveClientConfig(rawConfig, ctx, &clientcmd.ConfigOverrides{}, configAccess)
	restCfg, err := cfg.ClientConfig()
	if err != nil {
		err := fmt.Errorf("error when getting rest config: %s", err)
		k.Logger.Error(err)
		return err
	}

	clienset, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		err := fmt.Errorf("error when creating clientset err: %s", err)
		k.Logger.Error(err)
		return err
	}

	k.ClientSet = clienset
	k.Config = cfg
	k.Logger.Infof("context switched to: %s", ctx)

	return nil
}

// Will return list of available context
func (k *KubernetesController) CheckAvailableContext(ctx string) bool {
	rawConfig, err := k.Config.RawConfig()
	if err != nil {
		err := fmt.Errorf("error %s, getting starting config", err.Error())
		k.Logger.Error(err)
		return false
	}

	_, ok := rawConfig.Clusters[ctx]
	return ok
}

// Labeling nodes with given nodeName, key, and value
func (k *KubernetesController) LabelNodes(ctx context.Context, nodeName, key, val string) error {
	patch := []PatchObject{{
		Op:    PatchReplaceOP,
		Path:  "/metadata/labels/" + key,
		Value: val,
	}}

	patchAsByte, err := json.Marshal(patch)
	if err != nil {
		err := fmt.Errorf("error when converting struct to byte: %s", err)
		k.Logger.Error(err)
		return err
	}

	res, err := k.ClientSet.CoreV1().Nodes().Patch(ctx, nodeName, types.JSONPatchType, patchAsByte, apimetav1.PatchOptions{})
	if err != nil {
		err := fmt.Errorf("error when labeling nodes: %s", err)
		k.Logger.Error(err)
		return err
	}

	k.Logger.Infof("success labeling node %s with %s=%s", res.Name, key, val)
	return nil
}

// TODO IMAGE DUMMY BUAT IOT
// smart campus raspi di beberapa titik sensornya beberapa
//

// Deploying to image to the nodes
func (k *KubernetesController) Deploy(ctx context.Context, params DeployParams) (*apiappsv1.Deployment, error) {
	deployClient := k.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &apiappsv1.Deployment{
		ObjectMeta: apimetav1.ObjectMeta{
			Name: params.Name,
		},
		Spec: apiappsv1.DeploymentSpec{
			Replicas: int32Ptr(params.Replica),
			Selector: &apimetav1.LabelSelector{
				MatchLabels: params.Labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: apimetav1.ObjectMeta{
					Labels: params.Labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  params.Name + "-container",
							Image: params.Image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
					NodeSelector: params.Targets,
				},
			},
		},
	}

	res, err := deployClient.Create(ctx, deployment, apimetav1.CreateOptions{})

	// TODO Handle error
	if kubeerror.IsNotFound(err) {
		fmt.Println("INI ERR GAKETEMU YA ADIK")
	}

	if err != nil {
		return res, err
	}

	return res, nil
}

// List all deployments
func (k *KubernetesController) Get(ctx context.Context, params DeployParams) (*apiappsv1.Deployment, error) {
	deploymentsClient := k.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	result, err := deploymentsClient.Get(ctx, params.Name, apimetav1.GetOptions{})
	if err != nil {
		return result, err
	}

	return result, nil
}

// Update current deployments
func (k *KubernetesController) Patch(ctx context.Context, params DeployParams) {
	deploymentsClient := k.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(ctx, "demo-deployment", apimetav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		// TODO Logic update

		result.Spec.Replicas = int32Ptr(1)
		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13"
		_, updateErr := deploymentsClient.Update(ctx, result, apimetav1.UpdateOptions{})
		return updateErr
	})

	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
}

// Delete deployments
func (k *KubernetesController) Delete(ctx context.Context, params DeployParams) error {
	deploymentsClient := k.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	deletePolicy := apimetav1.DeletePropagationForeground
	err := deploymentsClient.Delete(ctx, params.Name, apimetav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	// TODO Handle Error

	if kubeerror.IsNotFound(err) {
		fmt.Println("INI ERROR NOT FOUND")
	}

	return err
}

// Check status cluster
func (k *KubernetesController) HealthCheck(ctx context.Context) error {
	podInterface := k.ClientSet.CoreV1().Pods(apiv1.NamespaceDefault)
	_, err := podInterface.List(ctx, apimetav1.ListOptions{})
	return err
}

func (k *KubernetesController) CheckDeploymentStatus(ctx context.Context, deploymentName string) bool {
	deployment, err := k.ClientSet.AppsV1().Deployments(apimetav1.NamespaceDefault).Get(ctx, deploymentName, apimetav1.GetOptions{})
	if err != nil {
		k.Logger.Errorf("deployment: error when getting deployment client err: %s", err)
		return false
	}

	if deployment.Status.AvailableReplicas == *deployment.Spec.Replicas {
		return true
	}

	return false
}
