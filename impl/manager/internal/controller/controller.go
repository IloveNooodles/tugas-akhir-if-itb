package controller

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type KubernetesController struct {
	Logger *logrus.Logger
	Config clientcmd.ClientConfig
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

	l.Infof("success connected to context: %s", rawCfg.CurrentContext)
	return &KubernetesController{
		Logger: l,
		Config: cfg,
	}
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

	k.Logger.Infof("current context %s", ctx)
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
	k.Config = cfg

	return nil
}

func (k *KubernetesController) CreateDeployment(replica int, label map[string]string) {}
