package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type downloadCmdResponse struct {
	uuid  string
	error string
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download file from storage",
	Long:  `Download file from storage`,
	Run: func(cmd *cobra.Command, args []string) {
		uuidFlag := cmd.Flag("uuid")
		pathFalg := cmd.Flag("PATH")
		isBackgroundCmd := cmd.Flag("background")

		if isBackgroundCmd.Value.String() == "true" {
			fmt.Println("Run download command in background")
		} else {
			fmt.Println("Run download command in interactive mode")
		}

		uuid := uuidFlag.Value.String()
		path := pathFalg.Value.String()

		uri := "/files/" + fmt.Sprintf("{%s}", uuid) + "?dirpath=" + path // GET
		fmt.Println(uri)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().BoolP("background", "b", false, "download file in background")
	downloadCmd.Flags().StringP("uuid", "u", "", "file uuid(required)")
	downloadCmd.Flags().StringP("PATH", "p", "", "path to file to save(required)")
	_ = downloadCmd.MarkFlagRequired("PATH")
	_ = downloadCmd.MarkFlagRequired("uuid")
}
