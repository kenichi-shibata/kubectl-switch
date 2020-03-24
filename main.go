package main

import (
	"github.com/kenichi-shibata/kubectl-switch/utils"
)

// Version change this to use go semver
var Version = "v0.0.2"

func main() {
	// cmd.Version = Version
	// cmd.Execute()
	// fmt.Println("underconst")
	utils.ParseKubectlVersion("v1.16.0-alpha.3")
}
