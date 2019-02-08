package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Create compressed artifact and upload to server",
	Long: `Creates a compressed artifact and uploads it to the artifacts server.

This command is the same as running ` + "`compress` and `upload`.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("archive called")
	},
}

func init() {
	rootCmd.AddCommand(archiveCmd)
	archiveCmd.Flags().StringP("target", "t", "./dist", "directory to archive")
}
