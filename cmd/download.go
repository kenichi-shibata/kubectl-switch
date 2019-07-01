package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/logutils"
	"github.com/kenichi-shibata/kubectl-switch/utils"
	"github.com/spf13/cobra"
)

// downloadCmd represents the downloa command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads kubectl binary",
	Long: `Downloads kubectl binary from prefix.
Stores the binary in ~/.kube/kubectl/ 
The version it will download will be from 
1. Command line flag either one: --stable --latest --kubectl-version
2. Environment variable: KUBECTL_VERSION
3. Config file: ~/.kube/kubectl/config
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the flags into variables start with log-level first
		logLevel, errLogLevel := cmd.Flags().GetString("log-level")
		// fmt.Printf("log-level:: %v", logLevel)
		if errLogLevel != nil {
			log.Print("[ERROR] ", errLogLevel)
			panic(errLogLevel)
		}
		switch logLevel {
		case "DEBUG":
		case "WARN":
		case "INFO":
		case "ERROR":
		default:
			log.Fatal("Unknown log level")
			panic(errors.New("Unknown Log Level"))
		}
		filter := &logutils.LevelFilter{
			Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
			MinLevel: logutils.LogLevel(logLevel), // this is where to set the log level
			Writer:   os.Stderr,
		}
		log.SetOutput(filter)

		versionFlag, errVersion := cmd.Flags().GetString("kubectl-version")
		if errVersion != nil {
			log.Print("[DEBUG] Could not get kubectl-version from flag trying in config")
		}
		configFlag, errConfig := cmd.Flags().GetString("config")
		if errConfig != nil {
			log.Print("[DEBUG] Did not find config File using default in ~/.kube/kubectl/config", errConfig)
		}

		log.Printf("[DEBUG] configFlag:: %v", configFlag)
		config := utils.ReadConfig(configFlag)
		kubectlVersion := utils.ParseKubectlVersion(versionFlag)

		var url, filepath string

		if len(kubectlVersion) > 0 {
			log.Print("[DEBUG] kubectlVersion exists ignoring config")
			kubectlConfig := config                       // copy the config parsed
			kubectlConfig.KubectlVersion = kubectlVersion // change the value of version
			url = utils.BuildURL(&kubectlConfig)
			filepath = utils.BuildFilepath(&kubectlConfig)

		} else {
			url = utils.BuildURL(&config)
			filepath = utils.BuildFilepath(&config)
		}

		filepathKubectl := utils.BuildFilepathKubectl()
		log.Printf("[INFO] downloading:: \n %v \n %v ...\n", filepath, url)
		err := downloadFile(filepath, url)
		if err != nil {
			log.Fatalf("Unable to download from %v to %v", url, filepath)
			panic(err)
		}
		errMod := os.Chmod(filepath, 0700)
		if errMod != nil {
			panic(errMod)
		}
		fmt.Println("\n##### Export the new path below or add it in bashrc or bash_profile to make it permanentt")
		fmt.Println("\nexport PATH=~/.kube/kubectl:$PATH\n")
		fmt.Println("kubectl version --client=true")

		errSoftlink := utils.SoftlinkKubectl(filepath, filepathKubectl)
		if errSoftlink != nil {
			panic(errSoftlink)
		}
	},
}

func init() {
	// stable := utils.StableVer()
	// latest := utils.LatestVer()

	prefix := utils.Prefix()
	kubectlVersion := utils.KubectlVersion()
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().BoolP("stable", "s", false, "use the stable version")
	downloadCmd.Flags().BoolP("latest", "l", false, "use the latest version")
	downloadCmd.Flags().StringP("prefix", "p", prefix, "Modify the prefix url where the binary will be downloaded from (Not needed most of the time)")
	downloadCmd.Flags().StringP("kubectl-version", "k", kubectlVersion, "Kubectl version to switch to")
}

// Download a file given a filepath where to save it and a url where the file exists assumes a single file
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
