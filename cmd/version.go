package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string
var buildDate string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of archivist",
	Long:  `All software has versions. This is archivist's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s (built %s)\n", version, buildDate)
	},
}
