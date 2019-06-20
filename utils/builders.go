package utils

import (
	"fmt"
	"log"
	"runtime"
)

var prefix string
var kubectlVersion string

// BuildURL builds the binary url where to download the binary and returns it
func BuildURL(config *Configuration) string {
	if config.KubectlPrefix != "" {
		prefix = config.KubectlPrefix
	} else {
		log.Fatalf("cannot get prefix from BuildURL")
	}

	if config.KubectlVersion != "" {
		kubectlVersion = config.KubectlVersion
	} else {
		log.Fatalf("cannot get kubectlVersion from BuildURL")
	}

	return fmt.Sprintf("%v/%v/bin/%v/%v/kubectl", prefix, kubectlVersion, runtime.GOOS, runtime.GOARCH)
}

// BuildFilepath Uses version() and home string to return the fullpath of the kubectl ex: ~/.kube/kubectl/kubectl-v1.14.3
func BuildFilepath(config *Configuration) string {
	if config.KubectlVersion != "" {
		kubectlVersion = config.KubectlVersion
	} else {
		log.Fatalf("cannot get kubectlVersion from BuildFilepath")
	}
	home, errCreateKubectlHome := CreateKubectlHome()
	if errCreateKubectlHome != nil {
		log.Fatalf("Unable to get the kubectl home")
		panic(errCreateKubectlHome)
	}
	return fmt.Sprintf("%v/kubectl-%v", home, kubectlVersion)
}

func BuildFilepathKubectl() string {
	home, errCreateKubectlHome := CreateKubectlHome()
	if errCreateKubectlHome != nil {
		log.Fatalf("Unable to get the kubectl home")
		panic(errCreateKubectlHome)
	}
	return fmt.Sprintf("%v/kubectl", home)
}