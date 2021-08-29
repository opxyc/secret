package cobra

import (
	"fmt"

	"github.com/opxyc/secret"
	"github.com/spf13/cobra"
)

var rmCmd = cobra.Command{
	Use:   "rm",
	Short: "Removes a key and it's value from secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Incorrect usage. Use help.")
			return
		}
		encodingKey, err := readKeyFromUser("key : ")
		if err != nil {
			fmt.Println("Failed to read the key.")
			return
		}

		verifyFilePath(&filePath)

		v := secret.New(encodingKey, filePath)
		key := args[0]
		err = v.Remove(key)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		fmt.Printf("removed '%s'!\n", key)
	},
}
