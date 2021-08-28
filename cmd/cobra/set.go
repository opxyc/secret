package cobra

import (
	"fmt"

	"github.com/opxyc/secret"
	"github.com/spf13/cobra"
)

var setCmd = cobra.Command{
	Use:     "set",
	Short:   "Sets a secret in your secret storage.",
	Example: "secret set key value -k encodingKey",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Incorrect usage. Use help.")
			return
		}
		encodingKey, err := getEncodingKey()
		if err != nil {
			fmt.Println("Failed to read the encoding key.")
			return
		}
		v := secret.File(encodingKey, filePath)
		key, value := args[0], args[1]
		err = v.Set(key, value)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Value set successfully!")
	},
}
