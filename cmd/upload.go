package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	viper.BindEnv("s3.bucket", "ARTIFACTS_S3_BUCKET")
	viper.BindEnv("s3.access-key", "ARTIFACTS_S3_ACCESS_KEY")
	viper.BindEnv("s3.secret-key", "ARTIFACTS_S3_SECRET_ACCESS_KEY")
}
