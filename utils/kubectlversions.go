package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type GithubKubernetesVersion struct {
	Version string `json:"tag_name"`
}

// method for data unmarshaled using GithubKubernetesVersion struct
func (g *GithubKubernetesVersion) String() string {
	return g.Version
}

// add a default timeout for http.Client
var httpClient = &http.Client{Timeout: 10 * time.Second}

// StableVer returns the stable kubectl version from github
func StableVer() string {
	resp, err := httpClient.Get("https://api.github.com/repos/kubernetes/kubernetes/releases/latest")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, ioerr := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(ioerr)
	}

	version := GithubKubernetesVersion{}
	jsonErr := json.Unmarshal(body, &version)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return version.String()
}

// LatestVer returns the latest kubectl verion from github (including the pre release)
func LatestVer() string {
	resp, err := httpClient.Get("https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, ioerr := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(ioerr)
	}

	// create array of the struct above and use that to unmarshal
	var arr []GithubKubernetesVersion
	jsonErr := json.Unmarshal(body, &arr)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// use the struct method String() defined above
	return arr[0].String()
}

// KubectlVersion returns the version either set by env var KUBECTL_VERSION
// or via the config file if it exists already
func KubectlVersion() string {
	version := os.Getenv("KUBECTL_VERSION")
	if version == "" {
		data := initializeConfigFile()
		return data.KubectlVersion
	}
	// TODO need to add input checker here
	// TODO get the stable version here https://storage.googleapis.com/kubernetes-release/release/stable.txt
	// TODO get the latest version here https://storage.googleapis.com/kubernetes-release/release/latest.txt
	// TODO return stable version by default
	return version
}

// WriteVersion writes the active version to the config file
// and writes all the downloaded files to the config file for list (local listing)
func WriteVersion() string {
	return ""
}
