package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

//  Struct for kubectl version and prefix
type Configuration struct {
	KubectlPrefix  string `json:"url_prefix"`
	KubectlVersion string `json:"version"`
}

// Seed the initial config file ~/.kube/kubectl/config
func seedData() Configuration {
	data := Configuration{
		KubectlPrefix:  "https://storage.googleapis.com/kubernetes-release/release",
		KubectlVersion: "v1.14.3",
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
	if prefix == "" {
		// check if ~/.kube/kubectl exists if not create it
		fmt.Println("creating ~/.kube/kubectl")
		home, err := createKubectlHome()
		if err != nil {
			panic(err)
		}
		// if env var does not exists try to read from ~/.kube/kubectl/config
		config := fmt.Sprintf("%v/config", home)
		if _, err := os.Stat(config); os.IsNotExist(err) {
			_, err := os.Create(config)
			if err != nil {
				panic(err)
			}
			seed, seedErr := json.MarshalIndent(seedData(), "", " ")
			if seedErr != nil {
				panic(seedErr)
			}
			fmt.Println("writing file ~/.kube/kubectl/config")
			noWriteErr := ioutil.WriteFile(config, seed, 0666)
			if noWriteErr != nil {
				panic(noWriteErr)
			}
		}
		fmt.Println("File ~/.kube/kubectl/config exists now reading it....")
		// config file is already exists at this point (created above or already exists) read it now
		configFile, err := ioutil.ReadFile(config)
		data := Configuration{}
		jsonErr := json.Unmarshal([]byte(configFile), &data)
		if jsonErr != nil {
			panic(jsonErr)
		}
		fmt.Println("Read URL Prefix: ", data.KubectlPrefix)
		fmt.Println("Read version: ", data.KubectlVersion)
		return data.KubectlPrefix
	}
	return "https://storage.googleapis.com/kubernetes-release/release"
}

// This function returns a string of the version
func version() string {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		versionArgs := os.Args[1]
		return versionArgs // uses cmd line arguments
	}
	// TODO need to add input checker here
	// TODO get the stable version here https://storage.googleapis.com/kubernetes-release/release/stable.txt
	// TODO get the latest version here https://storage.googleapis.com/kubernetes-release/release/latest.txt
	// TODO return stable version by default
	return "v1.14.3"
}

// Uses version() and home string to return the fullpath of the kubectl ex: ~/.kube/kubectl/kubectl-v1.14.3
func versionFile(home string) string {
	return fmt.Sprintf("%v/kubectl-%v", home, version())
}

// Returns the path of kubectl bin for softlinking string
func kubectlFile(home string) string {
	return fmt.Sprintf("%v/kubectl", home)
}

// Creates the kubectl home dir in ~/.kube/kubectl returns this directory following $HOME
func createKubectlHome() (string, error) {
	// create directory
	dir := fmt.Sprintf("%v/.kube/kubectl", os.Getenv("HOME"))
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}

// Builds the url where to download the binary from returns this url as string
// for example https://storage.googleapis.com/kubernetes-release/release/v1.14.0/bin/linux/amd64/kubectl
func buildURL() string {
	fmt.Printf("%v/%v/bin/%v/%v/kubectl", prefix(), version(), runtime.GOOS, runtime.GOARCH)
	return fmt.Sprintf("%v/%v/bin/%v/%v/kubectl", prefix(), version(), runtime.GOOS, runtime.GOARCH)
}

// Download a file given a filepath where to save it and a url where the file exists assumes a single file
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

// Wrapper for softlinking kubectl-vx.x.x to kubectl
func softlinkKubectl(oldname, newname string) error {
	if _, err := os.Lstat(newname); err == nil {
		if err := os.Remove(newname); err != nil {
			return fmt.Errorf("failed to unlink: %+v", err)
		}
	} else if os.IsNotExist(err) {
		return fmt.Errorf("failed to check symlink: %+v", err)
	}
	err := os.Symlink(oldname, newname)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	home, err := createKubectlHome()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n ====== downloading kubectl ver: %v from %v ...", version(), buildURL())
	if err := downloadFile(versionFile(home), buildURL()); err != nil {
		panic(err)
	}
	errMod := os.Chmod(versionFile(home), 0700)
	if errMod != nil {
		panic(errMod)
	}

	fmt.Println("\nexport PATH=~/.kube/kubectl:$PATH")
	fmt.Printf("Use kubectl-%v for execution", version())
	errSoftlink := softlinkKubectl(versionFile(home), kubectlFile(home))
	if errSoftlink != nil {
		panic(errSoftlink)
	}
	// TODO: make this into a cobra cmd line
	// TODO: Make version() read from the json file
	// TODO: Annotate this using godoc
	// TODO: create a list commmand to show which versions are available on the machine
	// TODO: create a list command to remotely show which versions are available
	// Sample structure {"versions" : ["v1.14.3","v1.14.0"], "version_active": "v1.14.3","url_prefix": ...}
}
