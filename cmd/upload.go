package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var uploadLog = log.WithFields(log.Fields{"stage": "upload"})

var uploadCmd = &cobra.Command{
	Use:   "upload [name of project] [path to file to upload]",
	Short: "upload an artifact to s3 artifacts bucket",
	Long:  "upload an artifact to s3 artifacts bucket",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := DoUpload(args)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	BindS3Flags(uploadCmd)
}

// DoUpload uploads the given asset to S3
func DoUpload(args []string) (err error) {
	name := args[0]
	assetPath := args[1]
	region := viper.GetString("s3.region")

	uploadLog.Info("checking connection info")

	if region == "" {
		return errors.New("must provide region--either set the env var S3_ARTIFACTS_REGION or use the --s3-region flag")
	}

	accessKey := viper.GetString("s3.access-key")

	if accessKey == "" {
		return errors.New("must provide aws access key--either set the env var AWS_ACCESS_KEY_ID or use the --s3-access-key flag")
	}

	secretKey := viper.GetString("s3.secret-key")

	if secretKey == "" {
		return errors.New("must provide aws secret access key--either set the env var AWS_SECRET_ACCESS_KEY or use the --s3-secret-access-key flag")
	}

	bucket := viper.GetString("s3.bucket")

	if bucket == "" {
		return errors.New("must provide a bucket name--either set the env var S3_ARTIFACTS_BUCKET or use the --s3-bucket flag")
	}

	uploadLog.Info("creating AWS session")
	s, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		return err
	}

	uploadLog.Info("getting branch name")
	branch := viper.GetString("git.branch")

	if branch == "" {
		gitCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
		out, err := gitCmd.CombinedOutput()

		if err != nil {
			return err
		}

		branch = string(out)
	}

	uploadLog.Info("determining destination key")
	branch = strings.Replace(strings.TrimSpace(branch), "/", "_", -1)
	assetName := filepath.Base(assetPath)
	key := strings.Join([]string{name, branch, assetName}, "/")
	err = addFileToS3(s, bucket, assetPath, key)

	if err != nil {
		return err
	}

	return nil
}

func addFileToS3(s *session.Session, bucket string, fileDir string, key string) error {
	file, err := os.Open(fileDir)

	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)

	file.Read(buffer)

	uploadLog.Info("uploading file to S3")
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(key),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}
