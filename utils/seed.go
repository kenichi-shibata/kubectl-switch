package utils

// SeedData seed the initial config file ~/.kube/kubectl/config
func SeedData() Configuration {
	data := Configuration{
		KubectlPrefix:  "https://storage.googleapis.com/kubernetes-release/release",
		KubectlVersion: "v1.14.3",
	}
	return data
}
