package main

import (
	"github.com/kenichi-shibata/kubectl-switch/cmd"
)

// Version change this to use go semver
var Version = "v0.0.1"

func main() {
	cmd.Version = Version
	cmd.Execute()
}
