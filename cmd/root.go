package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "archivist",
	Short: `Helper application to create artifact archives and upload them to
asset server`,
	Long: `archivist assists in creating artifact archives suitable for deployment
to an asset server. You can use it to generate compressed archives, upload an
archive to an asset server, or both in one command.`,
}

// Execute runs the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {}
