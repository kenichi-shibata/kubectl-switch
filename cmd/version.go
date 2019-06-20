package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

)
// Version kubectl-switch cli version
var Version string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kubectl-switch",
	Long:  `All software has versions. This is kubectl-switch's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
