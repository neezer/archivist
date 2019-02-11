package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// BindS3Flags binds s3 flags to the given command
func BindS3Flags(cmd *cobra.Command) {
	cmd.Flags().String("s3-bucket", "artifacts-kofile-systems", "name of bucket")
	cmd.Flags().String("s3-region", "us-west-2", "region for s3 bucket")
	cmd.Flags().String("s3-access-key", "", "access key for s3")
	cmd.Flags().String("s3-secret-access-key", "", "secret access key for s3")

	viper.BindEnv("s3.bucket", "ARTIFACTS_S3_BUCKET")
	viper.BindEnv("s3.region", "ARTIFACTS_S3_REGION")
	viper.BindEnv("s3.access-key", "ARTIFACTS_S3_ACCESS_KEY")
	viper.BindEnv("s3.secret-key", "ARTIFACTS_S3_SECRET_ACCESS_KEY")
	viper.BindPFlag("s3.bucket", cmd.Flags().Lookup("s3-bucket"))
	viper.BindPFlag("s3.region", cmd.Flags().Lookup("s3-region"))
	viper.BindPFlag("s3.access-key", cmd.Flags().Lookup("s3-access-key"))
	viper.BindPFlag("s3.secret-key", cmd.Flags().Lookup("s3-secret-access-key"))
}
