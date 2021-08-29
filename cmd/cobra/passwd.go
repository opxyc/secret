package cobra

import (
	"fmt"

	"github.com/opxyc/secret"
	"github.com/spf13/cobra"
)

var passwdCmd = cobra.Command{
	Use:   "passwd",
	Short: "Change current key of your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		currentEncodingKey, err := readKeyFromUser("key : ")
		if err != nil {
			fmt.Println("Failed to read the key.")
			return
		}

		newEncodingKey, err := readKeyFromUser("new key : ")
		if err != nil {
			fmt.Println("Failed to read the key.")
			return
		}

		verifyFilePath(&filePath)

		v := secret.New(currentEncodingKey, filePath)
		err = v.ChangeEncodingKey(newEncodingKey)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
		fmt.Println("key changed!")
	},
}
