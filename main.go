package main

// import "github.com/kenichi-shibata/kubectl-switch/utils"
import "github.com/kenichi-shibata/kubectl-switch/cmd"

// Version change this to use go semver
var Version = "v0.0.3"

func main() {
	cmd.Version = Version
	cmd.Execute()
	// utils.ParseKubectlVersion("v1.13.0")

}
