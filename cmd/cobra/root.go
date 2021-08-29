package cobra

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	filePath string
	rootCmd  = cobra.Command{
		Use:   "secret",
		Short: "Secret is an API key and other secrets manager",
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "the path to file where secrets are stored")
	rootCmd.AddCommand(&getCmd)
	rootCmd.AddCommand(&setCmd)
}

// getEncodingKey reads the encoding key from user
// with no echo to stdin.
func getEncodingKey() (string, error) {
	fmt.Print("encoding key : ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println()

	return strings.TrimSpace(string(bytePassword)), nil
}

// verifyFilePath checks if -f is set. If not, it will place
// $HOME/.secrets to filePaths variable.
func verifyFilePath(filePath *string) {
	if *filePath == "" {
		homeDir, _ := os.UserHomeDir()
		*filePath = filepath.Join(homeDir, ".secrets")
	}
}
