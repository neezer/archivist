package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress [dir to compress] [destination dir]",
	Short: "Compress an artifact directory",
	Long: `Compess a directory to a TarGZ archive. The resulting archive will be
named using the latest git SHA.

Please note the following:

  - requires git
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
	target := args[0]
	dest := args[1]

	target, err = filepath.Abs(target)

	if err != nil {
		return "", err
	}

	dest, err = filepath.Abs(dest)

	if err != nil {
		return "", err
	}

	targetInfo, err := os.Stat(target)

	if os.IsNotExist(err) {
		return "", err
	}

	if !targetInfo.IsDir() {
		return "", errors.New("target is not a directory")
	}

	err = os.MkdirAll(dest, os.ModePerm)

	if err != nil {
		return "", err
	}

	gitCmd := exec.Command("git", "rev-parse", "HEAD")
	out, err := gitCmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	sha := strings.TrimSpace(string(out))
	err = os.MkdirAll(sha, os.ModePerm)

	if err != nil {
		return "", err
	}

	removeTargetCopy := func() {
		os.RemoveAll(sha)
	}

	defer removeTargetCopy()

	err = copy.Copy(target, sha)

	if err != nil {
		return "", err
	}

	archiveFile = filepath.Join(dest, sha+".tar.gz")
	err = archiver.Archive([]string{sha}, archiveFile)

	if err != nil {
		os.RemoveAll(archiveFile)

		return "", err
	}

	return archiveFile, nil
}
