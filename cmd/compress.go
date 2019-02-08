package cmd

import (
	"fmt"
	"os"

	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
)

var target string
var dest string

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress an artifact directory",
	Long:  "Compess an artifact directory",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(target); os.IsNotExist(err) {
			fmt.Println(err)
			os.Exit(1)
			// log.Fatal("directory doesn't exist")
		}

		err := archiver.Archive([]string{target}, dest+".tar.gz")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
	compressCmd.Flags().StringVarP(&target, "target", "t", "./build", "directory to archive")
	compressCmd.Flags().StringVarP(&dest, "destination", "d", "./dist", "directory to save the archive to")
}
