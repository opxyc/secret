package cobra

import (
	"fmt"

	"github.com/opxyc/secret"
	"github.com/spf13/cobra"
)

var listCmd = cobra.Command{
	Use:   "list",
	Short: "Lists all keys stored",
	Run: func(cmd *cobra.Command, args []string) {
		encodingKey, err := readKeyFromUser("key : ")
		if err != nil {
			fmt.Println("Failed to read the encoding key.")
			return
		}

		verifyFilePath(&filePath)

		v := secret.New(encodingKey, filePath)
		keys, err := v.List()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
		if len(keys) == 0 {
			fmt.Println("It's empty!")
			return
		}
		fmt.Println("Stored keys\n---")
		for _, k := range keys {
			fmt.Println(k)
		}
	},
}
