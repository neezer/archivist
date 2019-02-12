package cmd

import (
	log "github.com/sirupsen/logrus"
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
		err := doArchive(args)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(archiveCmd)

	archiveCmd.Flags().StringVar(&dest, "dest", "./dist", "dir to place compressed artifact")
	BindS3Flags(archiveCmd)
}

func doArchive(args []string) error {
	name := args[0]
	target := args[1]

	log.Info("compressing target dir")
	archiveFile, err := DoCompress([]string{target, dest})

	if err != nil {
		return err
	}

	log.Info("processing upload")
	err = DoUpload([]string{name, archiveFile})

	if err != nil {
		return err
	}

	log.Info("complete")
	return nil
}
