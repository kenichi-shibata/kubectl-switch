package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/kenichi-shibata/kubectl-switch/utils"
	"github.com/hashicorp/logutils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubectl-switch",
	Short: "kubectl-switch a very fast kubectl-version switcher",
	Long: `Allows you to switch kubectl client version really quickly
Stores configuration in ~/.kube/kubectl/config and the binaries
in ~/.kube/kubectl/kubectl-*.
You will need to export this to PATH.`,
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

		// then get the rest of the flags
		// stable, errStable := cmd.Flags().GetBool("stable")
		// if errStable != nil {
		// 	log.Print("[ERROR] ", errStable)
		// 	panic(errStable)
		// }
		// latest, errLatest := cmd.Flags().GetBool("latest")
		// if errLatest != nil {
		// 	log.Print("[ERROR] ", errLatest)
		// 	panic(errLatest)
		// }
		// prefix, errPrefix := cmd.Flags().GetString("prefix")
		// if errPrefix != nil {
		// 	log.Print("[ERROR] ", errPrefix)
		// 	panic(errPrefix)
		// }
		// kubectlVersion, errKubectlVersion := cmd.Flags().GetString("kubectl-version")
		// if errKubectlVersion != nil {
		// 	log.Print("[ERROR] ", errKubectlVersion)
		// 	panic(errKubectlVersion)
		// }
		ConfigFlag, errConfig := cmd.Flags().GetString("config")
		if errConfig != nil {
			log.Print("[ERROR] ", errConfig)
			panic(errConfig)
		}

		log.Printf("[DEBUG] ARGS:: %v %v ", ConfigFlag, logLevel)
		// log.Printf("[DEBUG] ARGS:: %v %v %v %v %v %v ", stable, latest, prefix, kubectlVersion, ConfigFlag, logLevel)
		

	},
}

func init() {
	homedir := utils.Homedir()
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringP("config", "c", homedir + "/.kube/kubectl/config", "Where the config file is stored")
	rootCmd.PersistentFlags().StringP("log-level", "", "DEBUG", "The log level of the application [INFO|ERROR|WARN|DEBUG]")

}

// Execute stuff
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Print("[ERROR]", err)
		os.Exit(1)
	}
}
