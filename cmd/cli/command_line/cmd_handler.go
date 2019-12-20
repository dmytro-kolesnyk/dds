package command_line

import (
	"fmt"
	client "github.com/dmytro-kolesnyk/dds/cmd/cli/controller"
	"github.com/spf13/cobra"
	"os"
)

func Init(controller *client.Controller) {
	rootCmd := &cobra.Command{Long: "CLI utility for work with distribute data storage",}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Files list from storage",
		Long:  `Files list from storage`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Retrieving file list...")

			res, err := controller.ListCmdHandler("")
			if err != nil {
				fmt.Printf("-> failed to retrieve file list: %s", err)
				return
			}

			if res == nil {
				fmt.Printf("-> file list is empty")
				return
			}

			fmt.Println("-> file list is retrieved:")
			for i := range res.Files {
				fmt.Printf("[%d] - [%s] %s\n", i+1, res.Files[i].UUID, res.Files[i].FileName)
			}
		},
	}

	uploadCmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload file to storage",
		Long:  `Upload file to storage`,
		Run: func(cmd *cobra.Command, args []string) {

			var (
				strategyMap = map[string]string{"sc": "copy", "sf": "fragment", "sfc": "fragment_copy"}
				payload     = &client.UploadReq{}
			)

			strategy := cmd.Flag("strategy")
			storeLocal := cmd.Flag("local")
			path := cmd.Flag("PATH")
			payload.FilePath = path.Value.String()
			isBackgroundCmd := cmd.Flag("background")

			if isBackgroundCmd.Value.String() == "true" {               //TODO: Add loop for interactive mode
				fmt.Println("Run upload command in background")
			} else {
				fmt.Println("Run upload command in interactive mode")
			}

			if storeLocal.Value.String() == "true" {
				payload.StoreLocally = true
			} else {
				payload.StoreLocally = false
			}

			if strategy.Value.String() != "" {
				if val, ok := strategyMap[strategy.Value.String()]; ok {
					payload.Strategy = val
				} else {
					fmt.Printf("\"%s\" is a wrong strategy key, please run \"upload --help\" and choose correct value \n", strategy.Value.String())
					return
				}
			} else {
				fmt.Printf("Strategy not define, \"default\" value will be use\n")
				payload.Strategy = "default"
			}

			fmt.Println("Uploading file to storage...")

			res, err := controller.UploadCmdHandler(payload, "")
			if err != nil {
				fmt.Printf("-> failed to upload file to storage: %s", err)
				return
			}

			if res == nil {
				fmt.Printf("-> file uuid is empty")
				return
			}
			fmt.Println(*payload)
			fmt.Printf("uuid: [%s]\n", res.UUID)
		},
	}

	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "Download file from storage",
		Long:  `Download file from storage`,
		Run: func(cmd *cobra.Command, args []string) {
			uuidFlag := cmd.Flag("uuid")
			pathFalg := cmd.Flag("PATH")
			isBackgroundCmd := cmd.Flag("background")

			if isBackgroundCmd.Value.String() == "true" {              //TODO: Add loop for interactive mode
				fmt.Println("Run download command in background")
			} else {
				fmt.Println("Run download command in interactive mode")
			}

			uuid := uuidFlag.Value.String()
			path := pathFalg.Value.String()

			fmt.Println("Downloading file from storage...")

			uri := "/" + uuid + "?dirpath=" + path
			res, err := controller.DownloadCmdHandler(uri)
			if err != nil {
				fmt.Printf("-> failed to retrieve file from storage: %s", err)
				return
			}

			if res == nil {
				fmt.Printf("-> file uuid is empty")
				return
			}

			fmt.Printf("uuid: [%s]\n", res.UUID)
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete file from storage",
		Long:  `Delete file from storage`,
		Run: func(cmd *cobra.Command, args []string) {
			uuidFlag := cmd.Flag("uuid")
			isBackgroundCmd := cmd.Flag("background")

			if isBackgroundCmd.Value.String() == "true" {                       //TODO: Add loop for interactive mode
				fmt.Println("Run delete command in background")
			} else {
				fmt.Println("Run delete command in interactive mode")
			}

			uuid := uuidFlag.Value.String()

			uri := fmt.Sprintf("{%s}", uuid) //DELETE
			fmt.Println(uri)

			fmt.Println("Deleting file from storage...")

			res, err := controller.DeleteCmdHandler("/" + uuid)
			if err != nil {
				fmt.Printf("-> failed to delete file from storage: %s", err)
				return
			}

			if res == nil {
				fmt.Printf("-> file uuid is empty")
				return
			}

			fmt.Printf("uuid: [%s]\n", res.UUID)

		},
	}

	//TODO: Add Status command

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(deleteCmd)

	uploadCmd.Flags().StringP("strategy", "s", "", "save strategy:\n \tPossible key values:\n "+
		"\t\tsc - copy,\n \t\tsf - fragment,\n \t\tsfc - fragment-copy")
	uploadCmd.Flags().BoolP("background", "b", false, "save file in background")
	uploadCmd.Flags().BoolP("local", "l", false, "save file in local storage")
	uploadCmd.Flags().StringP("PATH", "p", "", "path to file to save (required)")
	_ = uploadCmd.MarkFlagRequired("PATH")

	downloadCmd.Flags().BoolP("background", "b", false, "download file in background")
	downloadCmd.Flags().StringP("uuid", "u", "", "file uuid(required)")
	downloadCmd.Flags().StringP("PATH", "p", "", "path to file to save(required)")
	_ = downloadCmd.MarkFlagRequired("PATH")
	_ = downloadCmd.MarkFlagRequired("uuid")

	deleteCmd.Flags().BoolP("background", "b", false, "delete file in background")
	deleteCmd.Flags().StringP("uuid", "u", "", "deleting file uuid(required)")
	_ = deleteCmd.MarkFlagRequired("uuid")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
