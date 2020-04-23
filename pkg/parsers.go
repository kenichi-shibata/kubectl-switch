package pkg

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

// Get all the supported version from a file in github https://raw.githubusercontent.com/kenichi-shibata/kubectl-switch/master/supported_versions this file is generated beforehand via a curl command
func ParseKubectlVersion(kubectlVersion string) (string, error) {
	resp, err := httpClient.Get("https://raw.githubusercontent.com/kenichi-shibata/kubectl-switch/master/supported_versions")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, ioerr := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(ioerr)
	}

	supported := false
	reader := bytes.NewReader(body)
	scanner := bufio.NewScanner(reader)
	// fmt.Println("it reached prescanner")
	for scanner.Scan() {
		// fmt.Println("===", kubectlVersion, "?", scanner.Text())
		if kubectlVersion == scanner.Text() {
			supported = true
		}
	}
	if scanner.Err() != nil {
		fmt.Printf(" > Failed!: %v\n", scanner.Err())
	}
	fmt.Println(kubectlVersion, "supported : ", supported)
	if supported {
		return kubectlVersion, nil
	} else {
		return "", errors.New("not supported")
	}
}
