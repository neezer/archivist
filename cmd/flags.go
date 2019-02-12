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

	viper.BindEnv("git.branch", "BRANCH_NAME")
	viper.BindEnv("git.sha", "GIT_COMMIT")
	viper.BindEnv("s3.bucket", "S3_ARTIFACTS_BUCKET")
	viper.BindEnv("s3.region", "S3_ARTIFACTS_REGION")
	viper.BindEnv("s3.access-key", "AWS_ACCESS_KEY_ID")
	viper.BindEnv("s3.secret-key", "AWS_SECRET_ACCESS_KEY")
	viper.BindPFlag("s3.bucket", cmd.Flags().Lookup("s3-bucket"))
	viper.BindPFlag("s3.region", cmd.Flags().Lookup("s3-region"))
	viper.BindPFlag("s3.access-key", cmd.Flags().Lookup("s3-access-key"))
	viper.BindPFlag("s3.secret-key", cmd.Flags().Lookup("s3-secret-access-key"))
}
