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
	appsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/retry"
)

type DeployParams struct {
	Replica int32
	Name    string
	Image   string
	Labels  map[string]string
	Targets map[string]string
}

type KubernetesController struct {
	Logger    *logrus.Logger
	Config    clientcmd.ClientConfig
	ClientSet *kubernetes.Clientset
}

func New(l *logrus.Logger) (*KubernetesController, error) {
	cfg, err := generateKubeConfig()
	if err != nil {
		err := fmt.Errorf("error when config path err: %s", err)
		l.Error(err)
		return nil, err
	}

	rawCfg, err := cfg.RawConfig()
	if err != nil {
		err := fmt.Errorf("error when reading raw cfg err: %s", err)
		l.Error(err)
		return nil, err
	}

	restCfg, err := cfg.ClientConfig()
	if err != nil {
		err := fmt.Errorf("error when reading rest cfg err: %s", err)
		l.Error(err)
		return nil, err
	}

	clienset, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		err := fmt.Errorf("error when creating clientset err: %s", err)
		l.Error(err)
		return nil, err
	}

	l.Infof("success connected to context: %s", rawCfg.CurrentContext)
	return &KubernetesController{
		Logger:    l,
		Config:    cfg,
		ClientSet: clienset,
	}, nil
}

func (k *KubernetesController) GetConfig() clientcmd.ClientConfig {
	return k.Config
}

func (k *KubernetesController) GetRawConfig() (api.Config, error) {
	rawCfg, err := k.Config.RawConfig()
	if err != nil {
		k.Logger.Errorf("error when getting raw config err: %s", err)
		return rawCfg, err
	}
	return rawCfg, nil
}

func (k *KubernetesController) GetRestConfig() (*rest.Config, error) {
	restCfg, err := k.Config.ClientConfig()
	if err != nil {
		k.Logger.Errorf("error when getting rest config err: %s", err)
		return restCfg, err
	}

	return restCfg, nil
}

func (k *KubernetesController) SwitchCluster(ctx string) error {
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

func (k *KubernetesController) GetNodeInterface() v1.NodeInterface {
	return k.ClientSet.CoreV1().Nodes()
}

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

func (k *KubernetesController) Get(ctx context.Context, params DeployParams) (*apiappsv1.Deployment, error) {
	deploymentsClient := k.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	result, err := deploymentsClient.Get(ctx, params.Name, apimetav1.GetOptions{})
	if err != nil {
		return result, err
	}

	return result, nil
}

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

func (k *KubernetesController) Apply(ctx context.Context, name, image string, labels map[string]string) (*apiappsv1.Deployment, error) {
	deployClient := k.ClientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	meta := metav1.ObjectMeta()
	meta.WithName(name)
	meta.WithLabels(labels)

	selector := metav1.LabelSelector()
	selector.WithMatchLabels(labels)

	ports := corev1.ContainerPort()
	ports.WithName("http")
	ports.WithProtocol(apiv1.ProtocolTCP)
	ports.WithContainerPort(80)

	containers := corev1.Container()
	containers.WithName(name)
	containers.WithImage(image)
	containers.WithPorts()

	spec := corev1.PodSpec()
	spec.WithContainers(containers)

	template := corev1.PodTemplateSpec()
	template.WithSpec(spec)
	template.ObjectMetaApplyConfiguration = meta

	deploymentSpec := appsv1.DeploymentSpec()
	deploymentSpec.WithReplicas(2)
	deploymentSpec.WithSelector(selector)
	deploymentSpec.WithTemplate(template)

	deploymentApply := appsv1.Deployment(name, apiv1.NamespaceDefault)
	deploymentApply.WithSpec(deploymentSpec)
	deploymentApply.ObjectMetaApplyConfiguration = meta

	res, err := deployClient.Apply(ctx, deploymentApply, apimetav1.ApplyOptions{})
	if err != nil {
		panic(err)
	}

	return res, err
}
