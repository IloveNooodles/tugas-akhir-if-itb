package controller

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

var (
	ErrHomePathNotFound = errors.New("home dir is not found")
)

// Generate kube config path from homedir
func getKubeConfigPath() (*string, error) {
	var kubeconfig *string
	home := homedir.HomeDir()

	if home == "" {
		return nil, ErrHomePathNotFound
	}

	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	return kubeconfig, nil
}

// Generate default kube config
//
// This will be used for the initial setup of kubernetes cluster
func generateKubeConfig() (clientcmd.ClientConfig, error) {
	path, err := getKubeConfigPath()
	if err != nil {
		log.Printf("path not found error: %s\n", err)
		return &clientcmd.DefaultClientConfig, err
	}

	cfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: *path},
		&clientcmd.ConfigOverrides{})

	return cfg, nil
}

// validate raw k8s config
func validateContext(config *clientcmdapi.Config, context string) error {
	if len(context) == 0 {
		return errors.New("empty context names are not allowed")
	}

	for name := range config.Contexts {
		if name == context {
			return nil
		}
	}

	return fmt.Errorf("no context exists with the name: %q", context)
}

func int32Ptr(i int32) *int32 { return &i }

func boolPtr(b bool) *bool { return &b }
