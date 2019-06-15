package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/logutils"
	"github.com/kenichi-shibata/kubectl-switch/utils"
	"github.com/spf13/cobra"
)

type configuration utils.Configuration

var rootCmd = &cobra.Command{
	Use:   "kubectl-switch",
	Short: "kubectl-switch a very fast kubectl-version switcher",
	Long: `Allows you to switch kubectl client version really quickly
Stores configuration in ~/.kube/kubectl/config and the binaries
in ~/.kube/kubectl/kubectl-*.
You will need to export this to PATH.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Do Stuff Here
	// 	fmt.Println("TEST")
	// },
}

func init() {
	// stable := utils.StableVer()
	// latest := utils.LatestVer()
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"), // this is where to set the log level
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	prefix := prefix()
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringP("prefix", "p", prefix, "Modify the prefix url where the binary will be downloaded from (Not needed most of the time)")
	rootCmd.PersistentFlags().StringP("kubectl-version", "k", "v1.14.3", "Kubectl version to switch to")
	rootCmd.PersistentFlags().StringP("config", "c", "$HOME/.kube/kubectl/config", "Where the config file is stored")
	rootCmd.PersistentFlags().StringP("log-level", "", "WARN", "The log level of the application")
	rootCmd.PersistentFlags().BoolP("stable", "s", false, "use the stable version")
	rootCmd.PersistentFlags().BoolP("latest", "l", false, "use the latest version")
}

// Execute stuff
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Print("[ERROR]", err)
		os.Exit(1)
	}
}

func prefix() string {
	//  check env var KUBECTL_PREFIX
	prefix := os.Getenv("KUBECTL_PREFIX")
	if prefix == "" {
		data := initializeConfigFile()
		return data.KubectlPrefix
	}
	return prefix
}

func initializeConfigFile() configuration {
	// check if ~/.kube/kubectl exists if not create it
	log.Print("[DEBUG] checking ~/.kube/kubectl")
	home, err := utils.CreateKubectlHome()
	if err != nil {
		panic(err)
	}
	// if env var does not exists try to read from ~/.kube/kubectl/config
	config := fmt.Sprintf("%v/config", home)
	if _, err := os.Stat(config); os.IsNotExist(err) {
		_, err := os.Create(config)
		if err != nil {
			panic(err)
		}
		seed, seedErr := json.MarshalIndent(utils.SeedData(), "", " ")
		if seedErr != nil {
			panic(seedErr)
		}
		log.Print("[DEBUG] writing file ~/.kube/kubectl/config")
		noWriteErr := ioutil.WriteFile(config, seed, 0666)
		if noWriteErr != nil {
			panic(noWriteErr)
		}
	}
	log.Print("[DEBUG] File ~/.kube/kubectl/config exists now reading it....")
	// config file is already exists at this point (created above or already exists) read it now
	configFile, err := ioutil.ReadFile(config)
	data := configuration{}
	jsonErr := json.Unmarshal([]byte(configFile), &data)
	if jsonErr != nil {
		panic(jsonErr)
	}
	return data
}
