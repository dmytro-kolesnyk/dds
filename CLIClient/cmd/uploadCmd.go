package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type uploadCmdResponse struct {
	uuid  string
	error string
}

var (
	//response map[string]string
	uploadRequest = make(map[string]string)
	strategyMap   = map[string]string{"sc": "copy", "sf": "fragment", "sfc": "fragment-copy"}

	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload file to storage",
		Long:  `Upload file to storage`,
		Run: func(cmd *cobra.Command, args []string) {
			strategy := cmd.Flag("strategy")
			storeLocal := cmd.Flag("local")
			path := cmd.Flag("PATH")
			uploadRequest["path"] = path.Value.String()
			isBackgroundCmd := cmd.Flag("background")

			if isBackgroundCmd.Value.String() == "true" {
				fmt.Println("Run upload command in background")
			} else {
				fmt.Println("Run upload command in interactive mode")
			}

			if storeLocal.Value.String() == "true" {
				uploadRequest["store_local"] = "true"
			} else {
				uploadRequest["store_local"] = "false"
			}

			if strategy.Value.String() != "" {
				if val, ok := strategyMap[strategy.Value.String()]; ok {
					uploadRequest["strategy"] = val
				} else {
					fmt.Printf("\"%s\" is a wrong strategy key, please run \"upload --help\" and choose correct value \n", strategy.Value.String())
				}
			} else {
				fmt.Printf("Strategy not define, \"default\" value will be use\n")
				uploadRequest["strategy"] = "default"
			}

			fmt.Println(uploadRequest)
			//req, _ := json.Marshal(uploadRequest)
			//fmt.Println(req)
			//err := json.Unmarshal(req, &response)
			//if err != nil{
			//	fmt.Println(err)
			//}
			//fmt.Println(response)
		},
	}
)

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringP("strategy", "s", "", "save strategy:\n \tPossible key values:\n "+
		"\t\tsc - copy,\n \t\tsf - fragment,\n \t\tsfc - fragment-copy")
	uploadCmd.Flags().BoolP("background", "b", false, "save file in background")
	uploadCmd.Flags().BoolP("local", "l", false, "save file in local storage")
	uploadCmd.Flags().StringP("PATH", "p", "", "path to file to save (required)")
	_ = uploadCmd.MarkFlagRequired("PATH")
}
