package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type deleteCmdResponse struct {
	uuid  string
	error string
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete file from storage",
	Long:  `Delete file from storage`,
	Run: func(cmd *cobra.Command, args []string) {
		uuidFlag := cmd.Flag("uuid")
		isBackgroundCmd := cmd.Flag("background")

		if isBackgroundCmd.Value.String() == "true" {
			fmt.Println("Run delete command in background")
		} else {
			fmt.Println("Run delete command in interactive mode")
		}

		uuid := uuidFlag.Value.String()

		uri := "/files/" + fmt.Sprintf("{%s}", uuid) //DELETE
		fmt.Println(uri)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("background", "b", "", "delete file in background")
	deleteCmd.Flags().StringP("uuid", "u", "", "deleting file uuid(required)")
	_ = deleteCmd.MarkFlagRequired("uuid")
}
