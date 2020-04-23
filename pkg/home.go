package pkg

import (
	"fmt"
	"os"
)

// CreateKubectlHome creates the kubectl home dir in ~/.kube/kubectl returns this directory following $HOME
func CreateKubectlHome() (string, error) {
	// create directory
	dir := fmt.Sprintf("%v/.kube/kubectl", os.Getenv("HOME"))
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}
