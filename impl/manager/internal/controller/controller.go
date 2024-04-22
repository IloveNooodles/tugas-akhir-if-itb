package controller

import (
	"fmt"

	"github.com/sirupsen/logrus"
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

func New(l *logrus.Logger) *KubernetesController {
	cfg, err := generateKubeConfig()
	if err != nil {
		l.Fatalf("error when config path err: %s", err)
	}

	rawCfg, err := cfg.RawConfig()
	if err != nil {
		l.Fatalf("error when reading raw cfg err: %s", err)
	}

	restCfg, err := cfg.ClientConfig()
	if err != nil {
		l.Fatalf("error when reading rest cfg err: %s", err)
	}

	clienset, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		l.Fatalf("error when creating clientset err: %s", err)
	}

	l.Infof("success connected to context: %s", rawCfg.CurrentContext)
	return &KubernetesController{
		Logger:    l,
		Config:    cfg,
		ClientSet: clienset,
	}
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

func (k *KubernetesController) GetNodeInterface() v1.NodeInterface {
	return k.ClientSet.CoreV1().Nodes()
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
