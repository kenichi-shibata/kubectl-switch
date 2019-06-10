package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

//  "Configuration": struct for kubectl version and prefix
type Configuration struct {
	kubectlPrefix  string `json: "prefix"`
	kubectlVersion string `json: "version"`
}

func seedData() *Configuration {
	data := &Configuration{
		kubectlPrefix:  "https://storage.googleapis.com/kubernetes-release/release",
		kubectlVersion: "v1.14.3",
	}
	return data
}

// Get prefix url downloading the binary
// Most of the time it will be "https://storage.googleapis.com/kubernetes-release/release"
// A file will be created in ~/.kube/kubectl/config if you do not have this env var KUBECTL_PREFIX set
// The file will have the url "https://storage.googleapis.com/kubernetes-release/release" by default
func prefix() string {
	//  check env var KUBECTL_PREFIX
	prefix := os.Getenv("KUBECTL_PREFIX")
	if prefix != "" {
		// check if ~/.kube/kubectl exists if not create it
		_, err := createKubectlHome()
		if err != nil {
			panic(err)
		}
		// if env var does not exists try to read from ~/.kube/kubectl/config
		// config := fmt.Sprintf("%v/config", home)
		// if _, err := os.Stat(config); os.IsNotExist(err) {
		// 	configFile, err := os.Create(config)
		// 	seed, _ := json.MarshalIndent(seedData(), "", " ")
		// err := ioutil.WriteFile(config, seed, 0666)
		// if err != nil {
		// 	panic(err)
		// }
		// }
		// else if the config file is already exists don't try to recreate it but read it instead
		// file, err := ioutil.ReadFile(fmt.Sprintf("%v/config"))
	}
	return "https://storage.googleapis.com/kubernetes-release/release"
}

func version() string {
	return "v1.14.3" // change this to input from user
}

func versionFile(dir string) string {
	return fmt.Sprintf("%v/kubectl-%v", dir, version())
}

func createKubectlHome() (string, error) {
	// create directory
	dir := fmt.Sprintf("%v/.kube/kubectl", os.Getenv("HOME"))
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func buildURL() string {
	return fmt.Sprintf("%v/%v/bin/%v/%v/kubectl", prefix(), version(), runtime.GOOS, runtime.GOARCH)
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	dir, err := createKubectlHome()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n ====== downloading kubectl ver: %v from %v", version(), buildURL())
	if err := downloadFile(versionFile(dir), buildURL()); err != nil {
		panic(err)
	}
}
