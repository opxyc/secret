package cobra

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var RootCmd = cobra.Command{
	Use:   "secret",
	Short: "Secret is an API key and other secrets manager",
}

var filePath string

func init() {
	RootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "the path to file where secrets are stored")
	RootCmd.AddCommand(&getCmd)
	RootCmd.AddCommand(&setCmd)
}

func getEncodingKey() (string, error) {
	fmt.Print("encoding key : ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println()

	return strings.TrimSpace(string(bytePassword)), nil
}
