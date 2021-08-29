package cobra

import (
	"fmt"

	"github.com/opxyc/secret"
	"github.com/spf13/cobra"
)

var getCmd = cobra.Command{
	Use:   "get",
	Short: "Gets a secret from your secret storage",
	Long:  "secret get <key> -k <encodingKey>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Incorrect usage. Use help.")
			return
		}
		encodingKey, err := getEncodingKey()
		if err != nil {
			fmt.Println("Failed to read the encoding key.")
			return
		}

		verifyFilePath(&filePath)

		v := secret.New(encodingKey, filePath)
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Printf("No value set for '%s'\n", key)
			return
		}
		fmt.Println(value)
	},
}
