package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// initializeConfigFile initializes the config file if it does not exists and then seeds it with SeedConfig
func initializeConfigFile() Configuration {
	// check if ~/.kube/kubectl exists if not create it
	// log.Print("[DEBUG] checking ~/.kube/kubectl")
	home, err := CreateKubectlHome()
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
		seed, seedErr := json.MarshalIndent(SeedData(), "", " ")
		if seedErr != nil {
			panic(seedErr)
		}
		log.Print("[DEBUG] writing file ~/.kube/kubectl/config")
		noWriteErr := ioutil.WriteFile(config, seed, 0666)
		if noWriteErr != nil {
			panic(noWriteErr)
		}
	}
	// log.Print("[DEBUG] File ~/.kube/kubectl/config exists now reading it....")
	// config file is already exists at this point (created above or already exists) read it now
	configFile, err := ioutil.ReadFile(config)
	data := Configuration{}
	jsonErr := json.Unmarshal([]byte(configFile), &data)
	if jsonErr != nil {
		panic(jsonErr)
	}
	return data
}
