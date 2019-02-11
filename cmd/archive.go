package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dest string

var archiveCmd = &cobra.Command{
	Args:  cobra.ExactArgs(2),
	Use:   "archive [name] [dir to archive]",
	Short: "Create compressed artifact and upload to server",
	Long: `Creates a compressed artifact and uploads it to the artifacts server.

This command is the same as running ` + "`compress` and `upload`.",
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(doArchive(args))
	},
}

func init() {
	rootCmd.AddCommand(archiveCmd)

	archiveCmd.Flags().StringVar(&dest, "dest", "./dist", "dir to place compressed artifact")
	BindS3Flags(archiveCmd)
}

func doArchive(args []string) int {
	name := args[0]
	target := args[1]
	archiveFile, err := DoCompress([]string{target, dest})

	if err != nil {
		fmt.Println(err)
		return 1
	}

	err = DoUpload([]string{name, archiveFile})

	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
