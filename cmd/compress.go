package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver"
	"github.com/otiai10/copy"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var compressCmd = &cobra.Command{
	Use:   "compress [dir to compress] [destination dir]",
	Short: "Compress an artifact directory",
	Long: `Compess a directory to a TarGZ archive. The resulting archive will be
named using the latest git SHA.

Please note the following:

  - requires git (if GIT_COMMIT is not already set)
  - CWD must be in the project root folder
  - program must have permissions to write to the CWD`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := DoCompress(args)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
}

// DoCompress compresses a target directory
func DoCompress(args []string) (archiveFile string, err error) {
	compressLog := log.WithFields(log.Fields{"stage": "compress"})
	target := args[0]
	dest := args[1]

	compressLog.Info("gathering info")

	target, err = filepath.Abs(target)

	if err != nil {
		return "", err
	}

	dest, err = filepath.Abs(dest)

	if err != nil {
		return "", err
	}

	compressLog.Info("checking target")

	targetInfo, err := os.Stat(target)

	if os.IsNotExist(err) {
		return "", err
	}

	if !targetInfo.IsDir() {
		return "", errors.New("target is not a directory")
	}

	compressLog.Info("making dest dir")
	err = os.MkdirAll(dest, os.ModePerm)

	if err != nil {
		return "", err
	}

	compressLog.Info("getting git SHA")
	sha := viper.GetString("git.sha")

	if sha == "" {
		gitCmd := exec.Command("git", "rev-parse", "HEAD")
		out, err := gitCmd.CombinedOutput()

		if err != nil {
			return "", err
		}

		sha = strings.TrimSpace(string(out))
	}

	shaDir := filepath.Join(dest, sha)

	compressLog.Info("making SHA directory")
	err = os.MkdirAll(shaDir, os.ModePerm)

	if err != nil {
		return "", err
	}

	removeTargetCopy := func() {
		compressLog.Info("cleaning up copied files")

		os.RemoveAll(sha)
	}

	defer removeTargetCopy()

	compressLog.Info("copying files from target to SHA directory")
	err = copy.Copy(target, shaDir)

	if err != nil {
		return "", err
	}

	files, err := ioutil.ReadDir(shaDir)

	if err != nil {
		return "", err
	}

	for _, file := range files {
		compressLog.Debug(file.Name())
	}

	archiveFile = filepath.Join(dest, sha+".tar.gz")

	compressLog.Info("removing existing archive, if present")
	os.RemoveAll(archiveFile)

	compressLog.Info("creating archive")
	err = archiver.Archive([]string{shaDir}, archiveFile)

	if err != nil {
		os.RemoveAll(archiveFile)

		return "", err
	}

	return archiveFile, nil
}
