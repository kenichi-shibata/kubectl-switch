package cmd

import (
	"log"
	"os"

	"github.com/kenichi-shibata/kubectl-switch/utils"
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
	},
}

func init() {
	homedir := utils.Homedir()
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringP("config", "c", homedir+"/.kube/kubectl/config", "Where the config file is stored")
	rootCmd.PersistentFlags().StringP("log-level", "", "INFO", "The log level of the application [INFO|ERROR|WARN|DEBUG]")
}

// Execute stuff
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Print("[ERROR]", err)
		os.Exit(1)
	}
}
