package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// SeedData seed the initial config file ~/.kube/kubectl/config
func SeedData() Configuration {
	data := Configuration{
		KubectlPrefix:  "https://storage.googleapis.com/kubernetes-release/release",
		KubectlVersion: "v1.14.3",
	}
	return data
}

// SeedSupportedVersions gets the first 100 releases in github and then stores them in ~/.kube/kubectl/supported_versions
func SeedSupportedVersions() string {
	resp, err := httpClient.Get("https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=100")
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

	for _, element := range arr {
		fmt.Println(element.String()) // create this as a file and pipe it to ~/.kube/kubectl/supported_versions
	}
	return ""
}
