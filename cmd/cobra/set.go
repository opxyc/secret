package cobra

import (
	"fmt"

	"github.com/opxyc/secret"
	"github.com/spf13/cobra"
)

var setCmd = cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage.",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Incorrect usage. Use help.")
			return
		}
		encodingKey, err := readKeyFromUser("key : ")
		if err != nil {
			fmt.Println("Failed to read the key.")
			return
		}

		verifyFilePath(&filePath)

		fmt.Printf("\rSetting..")
		v := secret.New(encodingKey, filePath)
		key, value := args[0], args[1]
		err = v.Set(key, value)
		if err != nil {
			fmt.Printf("\r%s\n", err.Error())
			return
		}
		fmt.Printf("\rValue set successfully!\n")
	},
}
