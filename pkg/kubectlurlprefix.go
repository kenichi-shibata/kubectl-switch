package pkg

import (
	"os"
)

// Prefix returns the kubectl url prefix where to get the binary most
// of the time its https://storage.googleapis.com/kubernetes-release/release
// But you can override it by setting KUBECTL_URL_PREFIX or changing it in the config file
func Prefix() string {
	//  check env var KUBECTL_PREFIX
	prefix := os.Getenv("KUBECTL_URL_PREFIX")
	if prefix == "" {
		data := initializeConfigFile()
		return data.KubectlPrefix
	}
	return prefix
}
