package utils

import "os"

// StableVer returns the stable kubectl version from the txt file
func StableVer() string {
	return "v1.14.3"
}

// LatestVer returns the latest kubectl verion from the latest.txt
func LatestVer() string {
	return "v1.16.0-alpha"
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
