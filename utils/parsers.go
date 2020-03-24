package utils

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
)

func ParseKubectlVersion(kubectlVersion string) (string, error) {
	// add a regex here for v<int>.<int>.<int>
	path, errAbs := filepath.Abs("supported_versions")
	if errAbs != nil {
		panic(errAbs)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	supported := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println("===", kubectlVersion, "?", scanner.Text())
		if kubectlVersion == scanner.Text() {
			supported = true
		}
	}
	// fmt.Println(kubectlVersion, "supported : ", supported)
	if supported {
		return kubectlVersion, nil
	} else {
		return "", errors.New("not supported")
	}
}
