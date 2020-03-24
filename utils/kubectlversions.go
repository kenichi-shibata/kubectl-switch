package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type githubKubernetesVersion struct {
	Version string `json:"tag_name"`
}

func (g *githubKubernetesVersion) String() string {
	return g.Version
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

// StableVer returns the stable kubectl version from the txt file
// TODO change this to fetch the stable version dynamically
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

	version := githubKubernetesVersion{}
	jsonErr := json.Unmarshal(body, &version)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(version.String())
	return version.String()
}

// LatestVer returns the latest kubectl verion from the latest.txt
// TODO change this to fetch the latest version dynamically
func LatestVer() string {
	return "v1.16.0-alpha.0"
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
