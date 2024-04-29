package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type KubernetesController struct {
	Logger    *logrus.Logger
	Config    clientcmd.ClientConfig
	ClientSet *kubernetes.Clientset
}

// Create new kubernetes Controller
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

// Return config instance
func (k *KubernetesController) GetConfig() clientcmd.ClientConfig {
	return k.Config
}

// Return raw config instance
func (k *KubernetesController) GetRawConfig() (api.Config, error) {
	rawCfg, err := k.Config.RawConfig()
	if err != nil {
		k.Logger.Errorf("error when getting raw config err: %s", err)
		return rawCfg, err
	}
	return rawCfg, nil
}

// Return rest config instnance
func (k *KubernetesController) GetRestConfig() (*rest.Config, error) {
	restCfg, err := k.Config.ClientConfig()
	if err != nil {
		k.Logger.Errorf("error when getting rest config err: %s", err)
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

// Labeling nodes with given nodeName, key, and value
func (k *KubernetesController) LabelNodes(ctx context.Context, nodeName, key, val string) error {
	patch := PatchObject{
		Op:    PatchAddOP,
		Path:  "/metadata/labels/" + key,
		Value: val,
	}

	patchAsByte, err := json.Marshal(patch)
	if err != nil {
		err := fmt.Errorf("error when converting struct to byte: %s", err)
		k.Logger.Error(err)
		return err
	}

	res, err := k.ClientSet.CoreV1().Nodes().Patch(ctx, nodeName, types.JSONPatchType, patchAsByte, metav1.PatchOptions{})
	if err != nil {
		err := fmt.Errorf("error when labeling nodes: %s", err)
		k.Logger.Error(err)
		return err
	}

	k.Logger.Infof("success labeling node %s with %s=%s", res.Name, key, val)
	return nil
}
