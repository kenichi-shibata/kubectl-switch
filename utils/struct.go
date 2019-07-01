package utils

// Configuration Struct for kubectl version and prefix
type Configuration struct {
	KubectlPrefix  string `json:"url_prefix"`
	KubectlVersion string `json:"version"`
}
